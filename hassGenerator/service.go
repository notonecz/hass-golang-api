package hassGenerator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/notonecz/hass-golang-api/rest"
)

func GenerateServiceFile(auth *rest.IMain) error {
	services, err := rest.GetServices(auth)
	if err != nil {
		return fmt.Errorf("failed to get services: %w", err)
	}

	servicesSlice, ok := services.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected services type: %T", services)
	}

	var builder strings.Builder

	builder.WriteString("package hass\n\n")
	builder.WriteString("import \"github.com/notonecz/hass-golang-api/rest\"\n\n")

	builder.WriteString("type Domain string\n\nconst (\n")

	usedDomainNames := make(map[string]int)
	domainConstNames := make(map[string]string)

	for _, domainItem := range servicesSlice {
		domainMap, ok := domainItem.(map[string]interface{})
		if !ok {
			continue
		}

		domainName, ok := domainMap["domain"].(string)
		if !ok {
			continue
		}

		constName := nameGen(upppreConver(domainName), usedDomainNames)
		domainConstNames[domainName] = constName
		builder.WriteString(fmt.Sprintf("\t%s Domain = %q\n", constName, domainName))
	}
	builder.WriteString(")\n\n")

	builder.WriteString("// Service is a marker interface for all service types\n")
	builder.WriteString("type Service interface {\n\tString() string\n}\n\n")

	for _, domainItem := range servicesSlice {
		domainMap, ok := domainItem.(map[string]interface{})
		if !ok {
			continue
		}

		domainName, ok := domainMap["domain"].(string)
		if !ok {
			continue
		}

		servicesMap, ok := domainMap["services"].(map[string]interface{})
		if !ok {
			continue
		}

		domainPrefix := upppreConver(domainName)
		typeName := domainPrefix + "Service"

		builder.WriteString(fmt.Sprintf("type %s string\n\n", typeName))
		builder.WriteString(fmt.Sprintf("func (s %s) String() string { return string(s) }\n\n", typeName))
		builder.WriteString("const (\n")

		usedServiceNames := make(map[string]int)

		for serviceName := range servicesMap {
			baseConstName := domainPrefix + upppreConver(serviceName)
			constName := nameGen(baseConstName, usedServiceNames)
			builder.WriteString(fmt.Sprintf("\t%s %s = %q\n", constName, typeName, serviceName))
		}
		builder.WriteString(")\n\n")

		funcName := "Post" + domainPrefix + "Service"
		builder.WriteString(fmt.Sprintf("func %s(auth *rest.IMain, service %s, payload string) (interface{}, error) {\n", funcName, typeName))
		builder.WriteString(fmt.Sprintf("\treturn rest.PostService(auth, string(%s), service.String(), payload)\n", domainConstNames[domainName]))
		builder.WriteString("}\n\n")
	}

	dirPath := "hass"
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filePath := filepath.Join(dirPath, "hass_services.go")
	if err := os.WriteFile(filePath, []byte(builder.String()), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Printf("Generátor hotov, soubor vytvořen: %s\n", filePath)
	return nil
}

func upppreConver(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			runes := []rune(part)
			runes[0] = unicode.ToUpper(runes[0])
			parts[i] = string(runes)
		}
	}
	return strings.Join(parts, "")
}

func nameGen(baseName string, used map[string]int) string {
	if _, exists := used[baseName]; !exists {
		used[baseName] = 1
		return baseName
	}

	counter := used[baseName]
	newName := fmt.Sprintf("%s%d", baseName, counter)

	for {
		if _, exists := used[newName]; !exists {
			used[baseName] = counter + 1
			used[newName] = 1
			return newName
		}
		counter++
		newName = fmt.Sprintf("%s%d", baseName, counter)
	}
}
