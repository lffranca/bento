package processor

import "github.com/lffranca/bento/processor/handler"

// Processor Processor
type Processor interface {
	Handler(handler.Request) (handler.Response, error)
}
