package hassGenerator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/notonecz/hass-golang-api/rest"
)

func generateServiceFile(auth *rest.IMain) error {

	fmt.Println("Generating service file...")

	services, err := rest.GetServices(auth)
	if err != nil {
		return err
	}

	servicesSlice, ok := services.([]interface{})
	if !ok {
		return err
	}

	var builder strings.Builder

	builder.WriteString("package " + auth.Id + "\n\n")
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
		typeName := "I" + domainPrefix + "Service"

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

		funcName := domainPrefix + "Service"

		builder.WriteString(fmt.Sprintf("func %s(auth *rest.IMain, service %s, payload interface{}) (interface{}, error) {\n", funcName, typeName))
		builder.WriteString(fmt.Sprintf("\treturn rest.PostService[interface{}](auth, string(%s), service.String(), payload)\n", domainConstNames[domainName]))
		builder.WriteString("}\n\n")

		builder.WriteString(fmt.Sprintf("func X%s(auth *rest.IMain, service %s, payload interface{}) interface{} {\n", funcName, typeName))
		builder.WriteString(fmt.Sprintf("\tcon, err := rest.PostService[interface{}](auth, string(%s), service.String(), payload)\n", domainConstNames[domainName]))
		builder.WriteString(fmt.Sprintf("\tif err != nil {panic(err)}\n"))
		builder.WriteString("\treturn con\n")
		builder.WriteString("}\n\n")

	}

	if err := os.MkdirAll(auth.Id, 0755); err != nil {
		return err
	}
	filePath := filepath.Join(auth.Id, "hass_services.go")
	if err := os.WriteFile(filePath, []byte(builder.String()), 0644); err != nil {
		return err
	}
	fmt.Printf("Service file generated in %s\n", filePath)
	return nil
}

func upppreConver(s string) string {
	return upperConverter(s, []string{"_"})
}

func upperConverter(s string, separators []string) string {
	for _, sep := range separators {
		s = strings.ReplaceAll(s, sep, " ")
	}

	parts := strings.Fields(s)

	for i, part := range parts {
		r := []rune(part)
		r[0] = unicode.ToUpper(r[0])
		parts[i] = string(r)
	}

	return strings.Join(parts, "")
}

func normalize(s string) string {

	s = strings.Replace(s, "_", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, ":", "", -1)
	s = strings.Replace(s, "+", "", -1)

	r := []rune(s)
	if unicode.IsDigit(r[0]) {
		s = "N" + s
	}

	return s
}

func selectEntityType(s string) (string, string) {
	parts := strings.Split(s, ".")
	return parts[0], parts[1]
}
