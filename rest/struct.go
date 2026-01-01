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
	Components            []string                  `json:"components"`
	ConfigDir             string                    `json:"config_dir"`
	Elevation             int64                     `json:"elevation"`
	Latitude              float64                   `json:"latitude"`
	LocationName          string                    `json:"location_name"`
	Longitude             float64                   `json:"longitude"`
	TimeZone              string                    `json:"time_zone"`
	UnitSystem            *IAPIComponentsUnitSystem `json:"unit_system"`
	Version               string                    `json:"version"`
	WhitelistExternalDirs []string                  `json:"whitelist_external_dirs"`
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
	Required  bool                 `json:"required,omitempty"`
	Example   any                  `json:"example,omitempty"`
	Filter    *IAPIFilter          `json:"filter,omitempty"`
	Selector  *IAPISelector        `json:"selector,omitempty"`
	Collapsed bool                 `json:"collapsed,omitempty"`
	Fields    map[string]IAPIField `json:"fields,omitempty"`
}

type IAPIFilter struct {
	SupportedFeatures []int            `json:"supported_features,omitempty"`
	Attribute         map[string][]any `json:"attribute,omitempty"`
}

type IAPISelector struct {
	Object            interface{}         `json:"object,omitempty"`
	Text              interface{}         `json:"text,omitempty"`
	Entity            interface{}         `json:"entity,omitempty"`
	Number            interface{}         `json:"number,omitempty"`
	ConfigEntry       interface{}         `json:"config_entry,omitempty"`
	Select            *IAPISelectorSelect `json:"select,omitempty"`
	Theme             interface{}         `json:"theme,omitempty"`
	Boolean           interface{}         `json:"boolean,omitempty"`
	Date              interface{}         `json:"date,omitempty"`
	Datetime          interface{}         `json:"datetime,omitempty"`
	Time              interface{}         `json:"time,omitempty"`
	Statistic         interface{}         `json:"statistic,omitempty"`
	BackupLocation    interface{}         `json:"backup_location,omitempty"`
	ConversationAgent interface{}         `json:"conversation_agent,omitempty"`
	ColorRgb          interface{}         `json:"color_rgb,omitempty"`
	ColorTemp         interface{}         `json:"color_temp,omitempty"`
	Constant          interface{}         `json:"constant,omitempty"`
	Icon              interface{}         `json:"icon,omitempty"`
	State             interface{}         `json:"state,omitempty"`
	Media             interface{}         `json:"media,omitempty"`
}

type IAPISelectorSelect struct {
	TranslationKey string   `json:"translation_key,omitempty"`
	Options        []string `json:"options,omitempty"`
	Sort           bool     `json:"sort,omitempty"`
	Multiple       bool     `json:"multiple,omitempty"`
	CustomValue    bool     `json:"custom_value,omitempty"`
}
