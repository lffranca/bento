package handler

// Request Request
type Request interface {
	GetPayload() []byte
	GetAttributes() map[string]interface{}
	GetAttribute(string) interface{}
}
