package entity

import (
	"time"
)

type Organization struct {
	ID        string    `json:"id" yaml:"id"`
	Name      string    `json:"name" yaml:"name"`
	Users     []User    `json:"users,omitempty" yaml:"users,omitempty"`
	Proxies   []Proxy   `json:"proxies,omitempty" yaml:"proxies,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}
