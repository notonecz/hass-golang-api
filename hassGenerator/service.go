package hassGenerator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/notonecz/hass-golang-api/rest"
)

func GenerateServiceFile(auth *rest.IMain, instance string) error {
	services, err := rest.GetServices(auth)
	if err != nil {
		return fmt.Errorf("failed to get services: %w", err)
	}

	servicesSlice, ok := services.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected services type: %T", services)
	}

	var builder strings.Builder

	builder.WriteString("package " + instance + "\n\n")
	builder.WriteString("import \"github.com/notonecz/hass-golang-api/rest\"\n\n")

	builder.WriteString("type Domain string\n\nconst (\n")

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

		constName := upppreConver(domainName)

		domainConstNames[domainName] = constName
		builder.WriteString(fmt.Sprintf("\t%s Domain = %q\n", constName, domainName))

	}

	builder.WriteString(")\n\n")
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

		for serviceName := range servicesMap {

			baseConstName := domainPrefix + upppreConver(serviceName)
			builder.WriteString(fmt.Sprintf("\t%s %s = %q\n", baseConstName, typeName, serviceName))

		}

		builder.WriteString(")\n\n")

		for serviceName, serviceRaw := range servicesMap {

			serviceMap, ok := serviceRaw.(map[string]interface{})
			if !ok {
				continue
			}

			payloadName := domainPrefix + upppreConver(serviceName) + "Payload"
			builder.WriteString(fmt.Sprintf("type %s struct {\n", payloadName))

			if fieldsRaw, ok := serviceMap["fields"].(map[string]interface{}); ok {
				for fieldName := range fieldsRaw {
					goField := upppreConver(fieldName)
					builder.WriteString(fmt.Sprintf(
						"\t%s interface{} `json:\"%s,omitempty\"`\n",
						goField,
						fieldName,
					))
				}
			}

			if _, ok := serviceMap["target"]; ok {
				builder.WriteString(
					"\tEntityID string `json:\"entity_id,omitempty\"`\n",
				)
			}

			builder.WriteString("}\n\n")
		}

		funcName := "Post" + domainPrefix + "Service"

		builder.WriteString(fmt.Sprintf("func %s(auth *rest.IMain, service %s, payload interface{}) (interface{}, error) {\n", funcName, typeName))
		builder.WriteString(fmt.Sprintf("\treturn rest.PostService(auth, string(%s), service.String(), payload)\n", domainConstNames[domainName]))
		builder.WriteString("}\n\n")

	}

	if err := os.MkdirAll(instance, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	filePath := filepath.Join(instance, "hass_services.go")
	if err := os.WriteFile(filePath, []byte(builder.String()), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	fmt.Printf("Generated file in %s\n", filePath)
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
