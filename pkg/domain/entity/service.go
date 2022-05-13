package entity

import (
	"time"
)

type Service struct {
	ID        string    `json:"id" yaml:"id"`
	Name      string    `json:"name" yaml:"name"`
	Host      string    `json:"host" yaml:"host"`
	Port      int       `json:"port" yaml:"port"`
	Protocol  string    `json:"protocol" yaml:"protocol"`
	URL       string    `json:"url,omitempty" yaml:"url,omitempty"`
	Tags      []string  `json:"tags,omitempty" yaml:"tags,omitempty"`
	Routes    []Route   `json:"routes,omitempty" yaml:"routes,omitempty"`
	Proxy     Proxy     `json:"proxy" yaml:"proxy"`
	CreatedAt time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}
