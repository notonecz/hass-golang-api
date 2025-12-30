package hassGenerator

import "github.com/notonecz/hass-golang-api/rest"

func Generate(auth *rest.IMain) error {

	err := generateServiceFile(auth)
	if err != nil {
		return err
	}

	err2 := generateEntityFolders(auth)
	if err2 != nil {
		return err2
	}

	return nil

}

func GenerateX(auth *rest.IMain) {

	err := generateServiceFile(auth)
	if err != nil {
		panic(err)
	}

	err2 := generateEntityFolders(auth)
	if err2 != nil {
		panic(err2)
	}

}
