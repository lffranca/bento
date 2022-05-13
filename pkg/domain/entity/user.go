package entity

import (
	"time"
)

type User struct {
	ID           string       `json:"id" yaml:"id"`
	Name         string       `json:"name" yaml:"name"`
	Email        string       `json:"email" yaml:"email"`
	Organization Organization `json:"organization,omitempty" yaml:"organization,omitempty"`
	CreatedAt    time.Time    `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt    time.Time    `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}
