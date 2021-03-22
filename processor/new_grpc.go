package processor

import "github.com/lffranca/bento/processor/grpc"

func NewGRPC() (Processor, error) {
	return &grpc.Processor{}, nil
}
