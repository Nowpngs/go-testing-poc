package modal

import modal "go-testing-poc/pkg/abstract"

type Status int

const (
	Pending Status = iota + 1
	Confirmed
	Cancelled
)

type Booking struct {
	modal.AbstractModal
	BookingNumber string
	Status        Status
	UserID        int
}
