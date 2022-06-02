package utils

import (
	"time"
)

func FromPointer[T comparable](value *T) (item T) {
	if value == nil {
		return
	}

	return *value
}

func FromArrayPointer[T comparable](values []*T) (items []T) {
	for _, value := range values {
		if value != nil {
			items = append(items, *value)
		}
	}

	return
}

func FromIntToTime(value *int) (item time.Time) {
	if value == nil {
		return
	}

	return time.UnixMilli(int64(*value))
}
