package rest

type Entity[T any] struct {
	EntityID   string `json:"entity_id"`
	State      string `json:"state"`
	Attributes T      `json:"attributes"`
}

type IAPI struct {
	Message string `json:"message"`
}

type IAPIConfigCheck struct {
	Errors string `json:"errors"`
	Result string `json:"result"`
}

type IAPIConfig struct {
	Components            []string                 `json:"components"`
	ConfigDir             string                   `json:"config_dir"`
	Elevation             int64                    `json:"elevation"`
	Latitude              float64                  `json:"latitude"`
	LocationName          string                   `json:"location_name"`
	Longitude             float64                  `json:"longitude"`
	TimeZone              string                   `json:"time_zone"`
	UnitSystem            IAPIComponentsUnitSystem `json:"unit_system"`
	Version               string                   `json:"version"`
	WhitelistExternalDirs []string                 `json:"whitelist_external_dirs"`
}

type IAPIComponentsUnitSystem struct {
	Length      string `json:"length"`
	Mass        string `json:"mass"`
	Temperature string `json:"temperature"`
	Volume      string `json:"volume"`
}

type IAPIDomain struct {
	Domain   string                 `json:"domain"`
	Services map[string]IAPIService `json:"services"`
}

type IAPIService struct {
	Fields map[string]IAPIField `json:"fields"`
}

type IAPIField struct {
	Required bool        `json:"required"`
	Example  any         `json:"example"`
	Selector interface{} `json:"selector"` // TODO: Add data types for selector
}
