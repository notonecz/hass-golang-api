package rest

func GetConfig(auth *IMain) (interface{}, error) {
	return comGet[interface{}](auth, "api/config")
}

func PostCheckConfig(auth *IMain) (interface{}, error) {
	return comPost[interface{}](auth, "api/config/core/check_config", "")
}

func GetApi(auth *IMain) (interface{}, error) {
	return comGet[interface{}](auth, "api/")
}

func GetComponents(auth *IMain) (interface{}, error) {
	return comGet[interface{}](auth, "api/components")
}
