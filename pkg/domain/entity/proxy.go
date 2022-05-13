package entity

import (
	"time"
)

type Proxy struct {
	ID           string       `json:"id" yaml:"id"`
	Name         string       `json:"name" yaml:"name"`
	Organization Organization `json:"organization,omitempty" yaml:"organization,omitempty"`
	Gateways     []Gateway    `json:"gateways,omitempty" yaml:"gateways,omitempty"`
	Routes       []Route      `json:"routes,omitempty" yaml:"routes,omitempty"`
	Services     []Service    `json:"services,omitempty" yaml:"services,omitempty"`
	CreatedAt    time.Time    `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt    time.Time    `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}
