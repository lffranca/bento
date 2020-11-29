package register

// ProcessorType ProcessorType
type ProcessorType int

const (
	// ProcessorHTTP ProcessorHTTP
	ProcessorHTTP ProcessorType = iota
	// ProcessorGRPC ProcessorGRPC
	ProcessorGRPC ProcessorType = iota
	// ProcessorSFTP ProcessorSFTP
	ProcessorSFTP ProcessorType = iota
	// ProcessorAWSS3 ProcessorAWSS3
	ProcessorAWSS3 ProcessorType = iota
	// ProcessorAWSSQS ProcessorAWSSQS
	ProcessorAWSSQS ProcessorType = iota
	// ProcessorRabbitMQ ProcessorRabbitMQ
	ProcessorRabbitMQ ProcessorType = iota
	// ProcessorEvaluateJSON ProcessorEvaluateJSON
	ProcessorEvaluateJSON ProcessorType = iota
)
