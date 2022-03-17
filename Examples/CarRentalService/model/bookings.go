package model

import (
	"time"
)

type Booking struct {
	Id                string
	UserId            string
	Created           time.Time
	BookingStartTime  time.Time
	BookingEndTime    time.Time
	VehicleType       VehicleType
	VehicleId         string // Vehicle specifics would be updated while handing over the vehicle to the person
	RCNumber          string
	Brand             string
	Company           string
	BookingAmount     string
	BookingPaymentIds []string // to account for multiple payments for a booking which would include advance and final payment mainly.
	Status            string
}

func (booking *Booking) SetBookingStatus(status string) {
	booking.Status = status
}

func (booking *Booking) GetBookingId() string {
	return booking.Id
}

type BookingService interface {
	GetAllBooking() ([]Booking, error)
	MarkBookingComplete(bookingId string)
	DeleteBooking(bookingId string)

	CreateBooking() (Booking, error)
	CancelBooking(bookingId string) (Booking, error) // Will

}

type BookingServiceImpl struct {
	Booking
}

func (bookingService *BookingServiceImpl) GetAllBooking() ([]Booking, error) {
	return BookingData, nil
}

func (bookingService *BookingServiceImpl) MarkBookingComplete(bookingId string) {
	for id, _ := range BookingData {
		if BookingData[id].GetBookingId() == bookingId {
			BookingData[id].SetBookingStatus("COMPLETE")
			return
		}
	}
}

func (bookingService *BookingServiceImpl) DeleteBooking(bookingId string) {

	vehicleService := VehicleServiceImpl{}
	vehicle := vehicleService.GetVehicleById(bookingService.VehicleId)
	vehicle.Release()

	newBookingList := []Booking{}
	for _, booking := range BookingData {
		if booking.Id == bookingId {
			continue
		}
		newBookingList = append(newBookingList, booking)
	}
	BookingData = newBookingList
}

func (bookingService *BookingServiceImpl) CreateBooking() (Booking, error) {
	// Update Vehicle Object
	vehicleService := VehicleServiceImpl{}
	vehicle := vehicleService.GetVehicleById(bookingService.VehicleId)
	vehicle.SetBookingId(bookingService.Id)
	vehicle.SetStatus(ON_RENT)
	vehicle.Save()

	BookingData = append(BookingData, bookingService.Booking)
	return bookingService.Booking, nil
}

func (bookingService *BookingServiceImpl) CancelBooking(bookingId string) (Booking, error) {

	vehicleService := VehicleServiceImpl{}
	vehicle := vehicleService.GetVehicleById(bookingService.VehicleId)
	vehicle.Release()

	for id, _ := range BookingData {
		if BookingData[id].GetBookingId() == bookingId {
			BookingData[id].Status = "CANCELLED"
			return BookingData[id], nil
		}
	}

	return Booking{}, nil
}

// Functionality Required
// - Make a booking : both admin and user
// - Record payment for the booking. This would be separate from the create booking flow.
// - Cancel the booking

type BookingRequest struct {
	UserId           string
	VehicleType      VehicleType
	VehicleBrand     string
	VehicleCompany   string
	BookingStartTime time.Time
	BookingEndTime   time.Time
}

type CancelBookingRequest struct {
	BookingId string
	UserId    string
}

func (req *BookingRequest) CreateBooking() (Booking, error) {

	// Search if the vehicle type is present in the inventory
	// - if yes, go ahead with the booking
	// - else return an error

	// Create a Booking, commit it to DB and return booking details

	return Booking{}, nil
}

func (req *CancelBookingRequest) CancelBooking() (Booking, error) {
	// Mark the booking as cancelled

	// Release a booked vehicle in inventory and increase the available count

	// Initiate any refund if the booking is eligible for the same.

	return Booking{}, nil
}

// Record payment for vehicle booking API
