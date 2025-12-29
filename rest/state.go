package rest

import "errors"

// DELETE

func DeleteState(auth *IMain, entity string) (interface{}, error) {
	return comDelete(auth, "api/states/"+entity)
}

// POST

func PostService(auth *IMain, domain string, service string, payload string) (interface{}, error) {
	return comPost(auth, "api/services/"+domain+"/"+service, payload)
}

func PostState(auth *IMain, entity string, payload string) (interface{}, error) {
	return comPost(auth, "api/states/"+entity, payload)
}

func PostEvent(auth *IMain, eventType string, payload string) (interface{}, error) {
	return comPost(auth, "api/events/"+eventType, payload)
}

func PostTemplate(auth *IMain, eventType string, payload string) (interface{}, error) {
	return comPost(auth, "api/template", payload)
}

// GET

func GetState(auth *IMain, entity string) (interface{}, error) {
	if entity != "" {
		return comGet(auth, "api/states/"+entity)
	} else {
		return nil, errors.New("no defined entity")
	}
}

func GetStates(auth *IMain) (interface{}, error) {
	return comGet(auth, "api/states")
}

func GetServices(auth *IMain) (interface{}, error) {
	return comGet(auth, "api/services")
}

func GetCalendar(auth *IMain, entity string) (interface{}, error) {
	return comGet(auth, "api/calendars")
}

func GetErrorLog(auth *IMain, entity string) (interface{}, error) {
	return comGet(auth, "api/error_log")
}

func GetEvents(auth *IMain, entity string) (interface{}, error) {
	return comGet(auth, "api/events")
}
