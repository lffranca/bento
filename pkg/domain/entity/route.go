package entity

import (
	"time"
)

type Route struct {
	ID        string              `json:"id" yaml:"id"`
	Name      string              `json:"name" yaml:"name"`
	Headers   map[string][]string `json:"headers,omitempty" yaml:"headers,omitempty"`
	Methods   []string            `json:"methods,omitempty" yaml:"methods,omitempty"`
	Paths     []string            `json:"paths,omitempty" yaml:"paths,omitempty"`
	Protocols []string            `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Tags      []string            `json:"tags,omitempty" yaml:"tags,omitempty"`
	Proxy     Proxy               `json:"proxy" yaml:"proxy"`
	Service   Service             `json:"service,omitempty" yaml:"service,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt time.Time           `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}
