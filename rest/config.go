package rest

func GetConfig(auth *IMain) (interface{}, error) {
	return comGet(auth, "api/config")
}

func PostCheckConfig(auth *IMain) (interface{}, error) {
	return comPost(auth, "api/config/core/check_config", "")
}

func GetApi(auth *IMain) (interface{}, error) {
	return comGet(auth, "api")
}

func GetComponents(auth *IMain) (interface{}, error) {
	return comGet(auth, "api/components")
}
