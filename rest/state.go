package rest

// DELETE

func DeleteState[T any](auth *IMain, entity string) (T, error) {
	return comDelete[T](auth, "api/states/"+entity)
}

// POST

func PostService[T any](auth *IMain, domain string, service string, payload interface{}) (T, error) {
	return IcomPost[T](auth, "api/services/"+domain+"/"+service, payload)
}

func PostServiceX[T any](auth *IMain, domain string, service string, payload interface{}) T {
	res, err := PostService[T](auth, domain, service, payload)
	if err != nil {
		panic(err)
	}
	return res
}

func PostState[T any](auth *IMain, entity string, payload string) (T, error) {
	return comPost[T](auth, "api/states/"+entity, payload)
}

func PostStateX[T any](auth *IMain, entity string, payload string) T {
	res, err := comPost[T](auth, "api/states/"+entity, payload)
	if err != nil {
		panic(err)
	}
	return res
}

func PostEvent[T any](auth *IMain, eventType string, payload string) (T, error) {
	return comPost[T](auth, "api/events/"+eventType, payload)
}

func PostIntentHandle[T any](auth *IMain, payload string) (T, error) {
	return comPost[T](auth, "api/intent/handle", payload)
}

func PostTemplate[T any](auth *IMain, payload string) (T, error) {
	return comPost[T](auth, "api/template", payload)
}

// GET

func GetState[T any](auth *IMain, entity string) (T, error) {
	return comGet[T](auth, "api/states/"+entity)
}

// GET

func GetStateX[T any](auth *IMain, entity string) T {
	res, err := GetState[T](auth, entity)
	if err != nil {
		panic(err)
	}
	return res
}

func GetStates(auth *IMain) ([]Entity[map[string]interface{}], error) {
	return comGet[[]Entity[map[string]interface{}]](auth, "api/states")
}

func GetServices(auth *IMain) ([]IAPIDomain, error) {
	return comGet[[]IAPIDomain](auth, "api/services")
}

func GetCalendars(auth *IMain) (interface{}, error) {
	return comGet[interface{}](auth, "api/calendars")
}

func GetErrorLog(auth *IMain) (interface{}, error) {
	return comGet[interface{}](auth, "api/error_log")
}

func GetEvents(auth *IMain) (interface{}, error) {
	return comGet[interface{}](auth, "api/events")
}
