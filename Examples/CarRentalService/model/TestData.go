package model

var UserData []UserAccount = []UserAccount{}
var AdminData []AdminAccount = []AdminAccount{}
var DriverData []DriverAccount = []DriverAccount{}

var VehicleData []Vehicle = []Vehicle{
	&Car{
		VehicleCommon: VehicleCommon{
			RCNumber: "1",
		},
		CurrentDriver: "",
		CNGEnabled:    false,
	},
}
var BookingData []Booking = []Booking{
	Booking{
		Id:        "1",
		VehicleId: "1",
	},
}
