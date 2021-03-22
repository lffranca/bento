package http

import (
	"context"
	"log"
	"net/http"

	"github.com/lffranca/bento/processor/handler"
)

type Processor struct {
	URL         string
	QueryParams map[string]string
	Headers     map[string]string
	Cookies     []*http.Cookie
}

func (item *Processor) Handler(context.Context, handler.Request) (handler.Response, error) {
	log.Println("HTTP HANDLER")
	return nil, nil
}
