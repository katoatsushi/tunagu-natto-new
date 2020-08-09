package models

import (
	"time"
)

type Style struct {
	ID        uint64    `json:"id"`
	StyleName string    `json:"style_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
