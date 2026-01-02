<img src="https://raw.githubusercontent.com/notonecz/hass-golang-api/refs/heads/main/assets/hass-beta0.5.png" alt="drawing"></img>
## Generate Service Files

```go
import (
	"your-proj/NAME_OF_INSTANCE"

	"github.com/notonecz/hass-golang-api/hassGenerator"
)

func main() {
    auth := rest.Init("NAME_OF_INSTANCE", "YOUR_TOKEN", "homeassistant.local:8123", false)
    
	hassGenerator.GenerateX(auth)
}
```

## Usage of generateed Service Files

> [!IMPORTANT]
> This only works if you have generated ServiceFiles (`hassGenerator.GenerateX(auth)`).
> If you don't want to use ServiceFiles look [Here](https://github.com/notonecz/hass-golang-api/blob/main/README.md#use-api-without-servicefiles).
### DOMAIN/SERIVCE api
```go
import (
	"github.com/notonecz/hass-golang-api/rest"
)

func main() {
	auth := rest.Init("NAME_OF_INSTANCE", "YOUR_TOKEN", "homeassistant.local:8123", false)

	DOMAIN.ServiceX( // X in ServiceX stands for auto panic function
		auth,
		DOMAIN.SERVICE, // SERVICE
		DOMAIN.[SERVICE]Payload{ // PAYLOAD
			EntityID: DOMAIN.ENTITYID,
		},
	)

	// LIGHT EXAMPLE

	Light.ServiceX(
		auth,
		Light.TurnOn,
		Light.TurnOnPayload{
			EntityID: Light.LivingRoomSpotlights,
		},
	)

}

```
### STATE API
```go
import (
	"github.com/notonecz/hass-golang-api/rest"
)

func main() {
	auth := rest.Init("NAME_OF_INSTANCE", "YOUR_TOKEN", "homeassistant.local:8123", false)

	// RETURN STRUCUTRE: type E(ENTITY) rest.Entity[(ENTITY)]
	// ENTTITY TYPE:
	type Entity[T any] struct {
		EntityID   string `json:"entity_id"`
		State      string `json:"state"`
		Attributes T      `json:"attributes"`
	}

	varName := DOMAIN.Get[ENTITY]X(auth).[DATA] // RETURNED DATA is a semi-dynamic structure ( E(Entity) )

	fmt.Println(light2)

	// LIGHT EXAMPLE

	// WITHOUT AUTOPANIC

	light, err := Light.GetLivingRoomSpotlights(auth)
	if err != nil {
		panic(err)
	}
	fmt.Println(light.State)
	// PRTINTS: off

	// WITH AUTO PANIC

	light2 := Light.GetLivingRoomSpotlightsX(auth).State
	fmt.Println(light2)
	// PRTINTS: off

	// EXAMPLE OF RETURN STRUCTURE: (RETURNS ELivingRoomSpotlights)

	type ELivingRoomSpotlights rest.Entity[SLivingRoomSpotlights]

	type Entity[T any] struct {
		EntityID   string `json:"entity_id"`
		State      string `json:"state"`
		Attributes T      `json:"attributes"`
	}

	type SLivingRoomSpotlights struct {
		SupportedColorModes []interface{} `json:"supported_color_modes"`
		ColorMode interface{} `json:"color_mode"`
		FriendlyName string `json:"friendly_name"`
		SupportedFeatures float64 `json:"supported_features"`
	}

	

}

```
## Use api without ServiceFiles
### DOMAIN/SERIVCE api
```go
import (
	"github.com/notonecz/hass-golang-api/rest"
)

func main() {
	auth := rest.Init("NAME_OF_INSTANCE", "YOUR_TOKEN", "homeassistant.local:8123", false)

	rest.PostServiceX[TYPE](
		auth,
		DOMAIN,
		SERVICE,
		map[string]any{
			PAYLOAD_KEY: PAYLOAD,
		},
	)

	// LIGHT EXAMPLE

	rest.PostServiceX[any]( // RETURN TYPE
		auth,
		"light", // DOMAIN
		"turn_on", // SERVICE
		map[string]any{ // PAYLOAD TYPE
			"entity_id": "light.living_room_spotlights",
		},
	)

}
```
### STATE API
```go
import (
	"github.com/notonecz/hass-golang-api/rest"
)

func main() {
	auth := rest.Init("NAME_OF_INSTANCE", "YOUR_TOKEN", "homeassistant.local:8123", false)

    state := rest.GetStateX[rest.Entity[any]](
        auth,
        "light.mini",
    )
    
    fmt.Println(state.State)
    // RETURNS: off
	
}
```

> [!NOTE]
> IN NEXT VERISON, PAYLOADING WILL BE MUCH EASIER THANKS TO DEFAULT STRUCTURES

## REST Functions

```go
func GetConfig(auth *IMain) (interface{}, error)
```
```go
func PostCheckConfig(auth *IMain) (interface{}, error)
```
```go
func GetApi(auth *IMain) (interface{}, error)
```
```go
func GetComponents(auth *IMain) (interface{}, error)
```
```go
func HelpServices(auth *IMain) // PRINTS Domains and Services
```
```go
func Generate(auth *rest.IMain) error // Generate structure of your home assistant instance
```
```go
func GenerateX(auth *rest.IMain) // SAME AS Generate but with auto panic
```
```go
func Init(id string, token string, ip string, secure bool) *IMain // GENERATE STRUCUTRE with: Id, token, ip, secure
```
```go
// REST FUNCTIONS, more on **https://developers.home-assistant.io/docs/api/rest**

// DELETE

func DeleteState[T any](auth *IMain, entity string) (T, error) {
return comDelete[T](auth, "api/states/"+entity)
}

// POST

func PostService[T any](auth *IMain, domain string, service string, payload interface{}) (T, error) {
return IcomPost[T](auth, "api/services/"+domain+"/"+service, payload)
}

func PostState[T any](auth *IMain, entity string, payload string) (T, error) {
return comPost[T](auth, "api/states/"+entity, payload)
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

func GetStates(auth *IMain) ([]Entity[map[string]interface{}], error) {
return comGet[[]Entity[map[string]interface{}]](auth, "api/states")
}

func GetServices(auth *IMain) (interface{}, error) {
return comGet[interface{}](auth, "api/services")
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
```

## About project
### Websocket support

I plan to add websocket support, currently only REST API is available
