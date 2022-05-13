package entity

import (
	"time"
)

type Gateway struct {
	ID        string      `json:"id" yaml:"id"`
	Name      string      `json:"name" yaml:"name"`
	Type      GatewayType `json:"type" yaml:"type"`
	Proxy     Proxy       `json:"proxy" yaml:"proxy"`
	CreatedAt time.Time   `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}
