package rest

import "fmt"

func HelpServices(auth *IMain) {
	srvInterface, err := GetServices(auth)
	if err != nil {
		panic(err)
	}

	servicesList, ok := srvInterface.([]any)
	if !ok {
		panic("expected []any from GetServices")
	}

	for _, domainItem := range servicesList {
		domainMap, ok := domainItem.(map[string]any)
		if !ok {
			continue
		}

		domainName, _ := domainMap["domain"].(string)
		fmt.Println("Domain:", domainName)

		services, ok := domainMap["services"].(map[string]any)
		if !ok {
			continue
		}

		for serviceName := range services {
			fmt.Println("-", serviceName)
		}
	}
}
