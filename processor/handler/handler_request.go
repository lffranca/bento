package handler

import "context"

// Request Request
type Request interface {
	GetContext() context.Context
	GetPayload() []byte
	GetOriginalPayload() []byte
	GetAttributes() map[string]interface{}
	GetAttribute(string) interface{}
}
