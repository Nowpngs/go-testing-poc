package service

import (
	"context"
	bookingModal "go-testing-poc/pkg/booking"
	bookingDto "go-testing-poc/pkg/booking/dto"
	bookingRepo "go-testing-poc/pkg/booking/repository"
)

type BookingService interface {
	CreateBooking(ctx context.Context, booking *bookingModal.Booking) error
	GetBookingWithUserList(ctx context.Context) ([]*bookingDto.BookingWithUserResponse, error)
}

type bookingService struct {
	bookingRepo bookingRepo.BookingRepository
}

func NewBookingService(bookingRepo bookingRepo.BookingRepository) BookingService {
	return &bookingService{bookingRepo: bookingRepo}
}

func (s *bookingService) CreateBooking(ctx context.Context, booking *bookingModal.Booking) error {
	return s.bookingRepo.CreateBooking(ctx, booking)
}

func (s *bookingService) GetBookingWithUserList(ctx context.Context) (
	[]*bookingDto.BookingWithUserResponse, error) {
	return s.bookingRepo.GetBookingWithUserList(ctx)
}
