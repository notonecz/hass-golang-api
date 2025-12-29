package rest

func GetState(auth *IMain, entity string) ([]byte, error) {
	return comGet(auth, "api/states/"+entity)
}

func PostService(auth *IMain, domain string, service string, payload string) ([]byte, error) {
	return comPost(auth, "api/services/"+domain+"/"+service, payload)
}
