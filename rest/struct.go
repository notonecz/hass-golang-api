package rest

type Entity[T any] struct {
	EntityID   string `json:"entity_id"`
	State      string `json:"state"`
	Attributes T      `json:"attributes"`
}
