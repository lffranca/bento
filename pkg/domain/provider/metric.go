package provider

import (
	"context"
)

type MetricProvider interface {
	CreateNewBusinessMetric(ctx context.Context)
}
