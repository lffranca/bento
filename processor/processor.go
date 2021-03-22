package processor

import (
	"context"

	"github.com/lffranca/bento/processor/handler"
)

// Processor Processor
type Processor interface {
	Handler(context.Context, handler.Request) (handler.Response, error)
}
