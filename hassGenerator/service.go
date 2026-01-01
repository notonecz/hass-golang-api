package hassGenerator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/notonecz/hass-golang-api/rest"
)

func generateServiceFiles(auth *rest.IMain) error {

	fmt.Println("Generating service file...")

	if err := os.RemoveAll(auth.Id); err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.MkdirAll(auth.Id, 0755); err != nil {
		return err
	}

	err := os.Chdir(auth.Id)
	if err != nil {
		return err
	}

	services, err := rest.GetServices(auth)
	if err != nil {
		return err
	}

	err = generateDomainFile(auth, services)
	if err != nil {
		return err
	}

	domainConstNames := make(map[string]string)

	for _, domainItem := range services {

		afields := false

		domainPrefix := upppreConver(domainItem.Domain)

		err = os.Mkdir(domainPrefix, 0755)
		if err != nil {
			return err
		}

		var builder strings.Builder

		builder.WriteString("package " + domainPrefix + "\n\n")
		builder.WriteString("import \"github.com/notonecz/hass-golang-api/rest\"\n\n")

		domainConstName := "C" + domainPrefix + "Domain"
		domainConstNames[domainItem.Domain] = domainConstName

		builder.WriteString("type Domain string\n\nconst (\n")

		builder.WriteString(fmt.Sprintf("\t%s Domain = \"%s\"\n)\n\n", domainConstName, domainItem.Domain))

		typeName := "I" + domainPrefix + "Service"

		builder.WriteString(fmt.Sprintf("type %s string\n\n", typeName))
		builder.WriteString(fmt.Sprintf("func (s %s) String() string { return string(s) }\n\n", typeName))
		builder.WriteString("const (\n")

		for serviceName := range domainItem.Services {
			baseConstName := upppreConver(serviceName)
			builder.WriteString(fmt.Sprintf("\t%s %s = %q\n", baseConstName, typeName, serviceName))
		}

		builder.WriteString(")\n\n")

		generatedEnums := map[string]bool{}

		for serviceName, service := range domainItem.Services {

			payloadName := upppreConver(serviceName) + "Payload"
			builder.WriteString(fmt.Sprintf("type %s struct {\n", payloadName))

			for fieldName := range service.Fields {
				if fieldName == "entity_id" {
					continue
				} else if fieldName == "advanced_fields" {
					afields = true
					goField := upppreConver(fieldName)
					builder.WriteString(fmt.Sprintf(
						"\t%s %s `json:\"%s,omitempty\"`\n",
						goField,
						"*Advanced"+upppreConver(serviceName)+"Payload",
						fieldName,
					))
					continue
				}
				goField := upppreConver(fieldName)
				builder.WriteString(fmt.Sprintf(
					"\t%s interface{} `json:\"%s,omitempty\"`\n",
					goField,
					fieldName,
				))

			}

			builder.WriteString(
				"\tEntityID string `json:\"entity_id\"`\n",
			)
			builder.WriteString("}\n\n")

			if afields {

				var selectFields []string

				ApayloadName := "Advanced" + upppreConver(serviceName) + "Payload"
				builder.WriteString(fmt.Sprintf("type %s struct {\n", ApayloadName))

				for fieldName := range service.Fields["advanced_fields"].Fields {

					typeDef := "interface{}"

					goField := upppreConver(fieldName)

					if service.Fields["advanced_fields"].Fields[fieldName].Selector.Select != nil {
						typeDef = "*I" + goField + "Select"
						selectFields = append(selectFields, fieldName)
					}

					builder.WriteString(fmt.Sprintf(
						"\t%s %s `json:\"%s,omitempty\"`\n",
						goField,
						typeDef,
						fieldName,
					))

				}

				builder.WriteString("}\n\n")

				for _, selectFieldName := range selectFields {

					tName := "I" + upppreConver(selectFieldName) + "Select"

					if generatedEnums[tName] {
						continue
					}
					generatedEnums[tName] = true

					fmt.Println(generatedEnums)

					builder.WriteString(fmt.Sprintf("type %s string\n\n", tName))
					builder.WriteString(fmt.Sprintf("func (s %s) String() string { return string(s) }\n\n", tName))
					builder.WriteString("const (\n")

					for _, optionName := range service.Fields["advanced_fields"].Fields[selectFieldName].Selector.Select.Options {
						baseConstName := upppreConver(optionName)
						builder.WriteString(fmt.Sprintf("\t%s %s = %q\n", baseConstName, tName, optionName))
					}

					builder.WriteString(")\n\n")
				}

			}

		}

		funcName := "Service"

		builder.WriteString(fmt.Sprintf("func %s(auth *rest.IMain, service %s, payload interface{}) (interface{}, error) {\n", funcName, typeName))
		builder.WriteString(fmt.Sprintf("\treturn rest.PostService[interface{}](auth, string(%s), service.String(), payload)\n", domainConstNames[domainItem.Domain]))
		builder.WriteString("}\n\n")

		builder.WriteString(fmt.Sprintf("func %sX(auth *rest.IMain, service %s, payload interface{}) interface{} {\n", funcName, typeName))
		builder.WriteString(fmt.Sprintf("\treturn rest.PostServiceX[interface{}](auth, string(%s), service.String(), payload)\n", domainConstNames[domainItem.Domain]))
		builder.WriteString("}\n\n")

		filePath := filepath.Join(domainPrefix, "SERVICE.go")
		if err := os.WriteFile(filePath, []byte(builder.String()), 0755); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filePath, err)
		}
		fmt.Printf("Entity generated in %s\n", filePath)
	}

	err = os.Chdir("..")
	if err != nil {
		return err
	}
	return nil
}

func generateDomainFile(auth *rest.IMain, services []rest.IAPIDomain) error {
	fmt.Println("Generating main service file...")

	domainConstNames := make(map[string]string)

	var builder strings.Builder

	builder.WriteString("package " + auth.Id + "\n\n")

	builder.WriteString("type Domain string\n\nconst (\n")

	for _, domainItem := range services {
		constName := upppreConver(domainItem.Domain)
		domainConstNames[domainItem.Domain] = constName
		builder.WriteString(fmt.Sprintf("\t%s Domain = %q\n", constName, domainItem.Domain))
	}

	builder.WriteString(")\n\n")
	builder.WriteString("type Service interface {\n\tString() string\n}\n\n")

	if err := os.WriteFile("domains.go", []byte(builder.String()), 0644); err != nil {
		return err
	}

	fmt.Printf("Service file generated: %s\n", "domains.go")

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
