package http

import (
	"context"
	"log"

	"github.com/lffranca/bento/processor/handler"
)

type Processor struct{}

func (item *Processor) Handler(context.Context, handler.Request) (handler.Response, error) {
	log.Println("HTTP HANDLER")
	return nil, nil
}
