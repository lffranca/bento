package processor

import "github.com/lffranca/bento/processor/http"

func NewHTTP() (Processor, error) {
	return &http.Processor{}, nil
}
