package grpc

import (
	"context"
	"log"

	"github.com/lffranca/bento/processor/handler"
)

type Processor struct{}

func (item *Processor) Handler(context.Context, handler.Request) (handler.Response, error) {
	log.Println("GRPC HANDLER")
	return nil, nil
}
