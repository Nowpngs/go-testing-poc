package repository

import (
	"context"
	"database/sql"
	bookingModal "go-testing-poc/pkg/booking"
	bookingDto "go-testing-poc/pkg/booking/dto"
	userDto "go-testing-poc/pkg/user/dto"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, booking *bookingModal.Booking) error
	GetBookingWithUserList(ctx context.Context) ([]*bookingDto.BookingWithUserResponse, error)
}

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) CreateBooking(ctx context.Context, booking *bookingModal.Booking) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO bookings (booking_number, status, user_id) VALUES ($1, $2, $3)", booking.BookingNumber, booking.Status, booking.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *bookingRepository) GetBookingWithUserList(ctx context.Context) ([]*bookingDto.BookingWithUserResponse, error) {
	query := `
		SELECT b.id, b.booking_number, b.status, u.id, u.username, u.email, u.role 
		FROM bookings b 
		JOIN users u ON b.user_id = u.id
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookings := make([]*bookingDto.BookingWithUserResponse, 0)
	for rows.Next() {
		booking := new(bookingDto.BookingWithUserResponse)
		user := new(userDto.UserResponse)
		err := rows.Scan(&booking.ID, &booking.BookingNumber, &booking.Status, &user.ID, &user.Username, &user.Email, &user.Role)
		if err != nil {
			return nil, err
		}
		booking.User = *user
		bookings = append(bookings, booking)
	}

	return bookings, nil
}
