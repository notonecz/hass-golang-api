package hassGenerator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/notonecz/hass-golang-api/rest"
)

func GenerateEntityFolders(auth *rest.IMain) error {
	if err := os.Chdir(auth.Id); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", auth.Id, err)
	}

	entities, err := rest.GetStates(auth)
	if err != nil {
		return fmt.Errorf("failed to get states: %w", err)
	}

	for _, entityItem := range entities {
		entityType, entitySelector := selectEntityType(entityItem.EntityID)
		entityTypeConverted := upppreConver(entityType)
		entitySelectorConverted := upppreConver(entitySelector)

		if err := os.MkdirAll(entityTypeConverted, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", entityTypeConverted, err)
		}

		var builder strings.Builder
		builder.WriteString("package " + entityTypeConverted + "\n\n")
		builder.WriteString("import \"github.com/notonecz/hass-golang-api/rest\"\n\n")
		builder.WriteString("type S" + entitySelectorConverted + " struct {\n")

		for key, value := range entityItem.Attributes {
			goType := mapJSONTypeToGo(value)
			builder.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n",
				normalize(
					upperConverter(key, []string{"_", " "})),
				goType, key),
			)
		}

		builder.WriteString("}\n\n")

		builder.WriteString("const ")
		builder.WriteString(fmt.Sprintf("%s = \"%s.%s\"\n", normalize(entitySelectorConverted), entityType, entitySelector))

		builder.WriteString("\n")
		builder.WriteString("type E" + entitySelectorConverted + " rest.Entity[S" + entitySelectorConverted + "]\n\n")
		builder.WriteString("func Get" + entitySelectorConverted + "(auth *rest.IMain) (E" + entitySelectorConverted + ", error) {\n")
		builder.WriteString("\t return rest.GetState[E" + entitySelectorConverted + "](auth, " + normalize(entitySelectorConverted) + ") \n}\n\n")

		builder.WriteString("func Get" + entitySelectorConverted + "X(auth *rest.IMain) E" + entitySelectorConverted + " {\n")
		builder.WriteString("\t con, err := rest.GetState[E" + entitySelectorConverted + "](auth, " + normalize(entitySelectorConverted) + ") \n")
		builder.WriteString("\t if err != nil {panic(err)}\n")
		builder.WriteString("\t return con\n}\n")

		filePath := filepath.Join(entityTypeConverted, entitySelectorConverted+".go")
		if err := os.WriteFile(filePath, []byte(builder.String()), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filePath, err)
		}
		fmt.Printf("Entity generated in %s\n", filePath)
	}

	return nil
}

func mapJSONTypeToGo(value interface{}) string {
	switch value.(type) {
	case string:
		return "string"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case map[string]interface{}:
		return "map[string]interface{}"
	case []interface{}:
		return "[]interface{}"
	default:
		return "interface{}"
	}
}
