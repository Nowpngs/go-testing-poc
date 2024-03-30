package modal

import "time"

type AbstractModal struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Valid bool
}