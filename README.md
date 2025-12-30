
# Home Assistant Golang API

## Generate hass-services

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


## Usage of generate hass-service functions

```go
import (
	"github.com/notonecz/hass-golang-api/rest"
)

func main() {
	auth := rest.Init("NAME_OF_INSTANCE", "YOUR_TOKEN", "homeassistant.local:8123", false)

	NAME_OF_INSTANCE.LightService( // DOMAIN
		auth,
		NAME_OF_INSTANCE.LightTurnOn, // SERVICE
		NAME_OF_INSTANCE.LightTurnOnPayload{ // PAYLOAD
			EntityID: Light.Mini,
		},
	)
}

```

`
hass.[Domain]Service(auth, NAME_OF_INSTANCE.[Domain][Service], NAME_OF_INSTANCE.[Domain][Service]Payload{ /* FIELDS */ })
`
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
