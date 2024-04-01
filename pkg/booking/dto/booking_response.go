package dto

import "go-testing-poc/pkg/user/dto"

type BookingWithUserResponse struct {
	ID            int              `json:"id"`
	BookingNumber string           `json:"booking_number"`
	Status        string           `json:"status"`
	User          dto.UserResponse `json:"user"`
}
