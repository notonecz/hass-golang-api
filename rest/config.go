package rest

func GetConfig(auth *IMain) (IAPIConfig, error) {
	return comGet[IAPIConfig](auth, "api/config")
}

func PostCheckConfig(auth *IMain) (IAPIConfigCheck, error) {
	return comPost[IAPIConfigCheck](auth, "api/config/core/check_config", "")
}

func GetApi(auth *IMain) (IAPI, error) {
	return comGet[IAPI](auth, "api/")
}

func GetComponents(auth *IMain) ([]string, error) {
	return comGet[[]string](auth, "api/components")
}
