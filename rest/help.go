package rest

import "fmt"

func HelpServices(auth *IMain) {
	srvInterface, err := GetServices(auth)
	if err != nil {
		panic(err)
	}

	for _, domainItem := range srvInterface {
		fmt.Println("Domain:", domainItem.Domain)

		for serviceName := range domainItem.Services {
			fmt.Println("-", serviceName)
			for fieldName := range serviceName {
				fmt.Println("-->", fieldName)
			}
		}
	}
}
