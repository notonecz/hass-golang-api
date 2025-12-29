package rest

func GetConfig(auth *IMain) ([]byte, error) {
	return comGet(auth, "api/config")
}
