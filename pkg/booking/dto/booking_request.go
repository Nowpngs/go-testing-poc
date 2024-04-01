package dto

import (
	bookingModal "go-testing-poc/pkg/booking"
)

type BookingCreateRequest struct {
	BookingNumber string              `json:"bookingNumber" binding:"required"`
	Status        bookingModal.Status `json:"status" binding:"required"`
	UserID        int                 `json:"userId" binding:"required"`
}
