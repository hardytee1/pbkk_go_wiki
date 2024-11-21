package models

import (
	"time"
)

type Blog struct {
	ID           uint           // Standard field for the primary key
	Title         string         // A regular string field
	Body         string         // A regular string field
	CreatedAt    time.Time      // Standard field for the creation time
	UpdatedAt    time.Time      // Standard field for the update time
}