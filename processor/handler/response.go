package handler

// Response Response
type Response interface {
	GetPayload() []byte
	GetAttributes() map[string]interface{}
	GetAttribute(string) interface{}
}
