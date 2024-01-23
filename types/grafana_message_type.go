package types

// https://gin-gonic.com/docs/examples/binding-and-validation/
// TODO implement
type GrafanaEvent struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type"`
}
