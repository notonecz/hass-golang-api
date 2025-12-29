
# Home Assistant Golang API

## Generate hass-services

```go
import (
	"your-proj/hass"

	"github.com/notonecz/hass-golang-api/hassGenerator"
)

func main() {
    auth := rest.Init("NAME_OF_INSTANCE", "YOUR_TOKEN", "homeassistant.local:8123", false)
    
	hassGenerator.GenerateServiceFile(auth)
}
```


## Usage of generate hass-service functions

```go
import (
	"github.com/notonecz/hass-golang-api/rest"
)

func main() {
	auth := rest.Init("NAME_OF_INSTANCE", "YOUR_TOKEN", "homeassistant.local:8123", false)

	hass.PostLightService(auth, hass.LightToggle, `{"entity_id": "light.mini"}`)
}

```

`
hass.Post[Domain]Service(auth, hass.[Domain][Service], jsonPayload)
`
## REST Functions

```go
func DeleteState(auth *IMain, entity string) (interface{}, error)
```
```go
func PostService(auth *IMain, domain string, service string, payload string)
```
```go
func PostState(auth *IMain, entity string, payload string) (interface{}, error)
```
```go
func PostEvent(auth *IMain, eventType string, payload string) (interface{}, error)
```
```go
func PostIntentHandle(auth *IMain, payload string) (interface{}, error)
```
```go
func PostTemplate(auth *IMain, payload string) (interface{}, error) 
```
```go
func GetState(auth *IMain, entity string) (interface{}, error)
```
```go
func GetStates(auth *IMain) (interface{}, error)
```
```go
func GetEvents(auth *IMain) (interface{}, error)
```
```go
func GetErrorLog(auth *IMain) (interface{}, error) 
```
```go
func GetCalendars(auth *IMain) (interface{}, error)
```
```go
func GetServices(auth *IMain) (interface{}, error)
```
```go
func GetApi(auth *IMain) (interface{}, error)
```
```go
func HelpServices(auth *IMain)
```
```go
func PostCheckConfig(auth *IMain) (interface{}, error)
```
```go
func GetComponents(auth *IMain) (interface{}, error)
```
```go
func GetConfig(auth *IMain) (interface{}, error)
```

### MORE ON https://developers.home-assistant.io/docs/api/rest

## About project
### Websocket suppor

I plan to add websocket support, currently only REST API is available
