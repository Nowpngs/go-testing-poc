package booking

import (
	bookingModal "go-testing-poc/pkg/booking"
	bookingDto "go-testing-poc/pkg/booking/dto"
	bookingService "go-testing-poc/pkg/booking/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	bookingService bookingService.BookingService
}

func NewBookingHandler(bookingService bookingService.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: bookingService}
}

// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a new booking with the input payload
// @Tags Bookings
// @Accept  json
// @Param booking body bookingDto.BookingCreateRequest true "Booking object that needs to be created"
// @Produce  json
// @Success 201 {object} bookingModal.Booking
// @Router /booking [post]
func (h *BookingHandler) CreateBookingHandler(c *gin.Context) {
	var bookingRequest bookingDto.BookingCreateRequest
	if err := c.BindJSON(&bookingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBooking := bookingModal.Booking{
		BookingNumber: bookingRequest.BookingNumber,
		Status:        bookingRequest.Status,
		UserID:        bookingRequest.UserID,
	}

	err := h.bookingService.CreateBooking(c.Request.Context(), &newBooking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBooking)
}

// GetBookingWithUserList godoc
// @Summary Get booking list with user list
// @Description Get a list of bookings with user list
// @Produce  json
// @Tags Bookings
// @Success 200 {array} bookingDto.BookingWithUserResponse
// @Router /booking/user [get]
func (h *BookingHandler) GetBookingWithUserListHandler(c *gin.Context) {
	bookings, err := h.bookingService.GetBookingWithUserList(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}
