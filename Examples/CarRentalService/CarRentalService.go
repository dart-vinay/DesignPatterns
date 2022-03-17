package CarRentalService

import (
	"DesignPatterns/Examples/CarRentalService/model"
	"github.com/labstack/gommon/log"
)

// Objectives
// - Admin Requirements:
// 		- would be able to add car to inventory.
// 		- Able to see all the bookings.
// 		-
// - User Side Requirements:
// 		- Able to book.
// 		- Able to pay for the booking
//  	- Able to see all his bookings
// 		- Cancel booking feature
// 		- Allow them to return the vehicle.

// DB requirements
// Booking Table, User details table.
// Vehicle Table : With all the details like RC number, engine number, registration number, brand, name, etc.
// Vehicle Inventory: By Vehicle Type. Would have details regarding : Total in ownership, available, on rent, others (gone for maintenance, etc)
// Vehicle rental history: To track the distance it has covered till now and other details, etc.

// Big challenges to look out
// - Avoid race condition while updating the inventory to keep it error free.
// -
func CarRentalService() {
	model.InitializeVehicleServices()

	bookingService := new(model.BookingServiceImpl)
	bookingService.Booking = model.Booking{
		Id:        "2",
		VehicleId: "1",
	}

	bookingService.CreateBooking()

	bookingService.CancelBooking("2")
	allBookings := model.BookingData
	vehicle := model.VehicleData[0].(*model.Car)
	log.Info(vehicle)
	log.Info(allBookings)

}
