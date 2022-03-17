package model

type VehicleType string
type STATUS string

// AddOns are initialized as singletons. They can be stored in db as well to provide flexibility of adding and modifying services without restarting server.
var (
	FueledTankService  *FueledTank
	AddOnDriverService *AddonDriver
	InsuranceService   *Insurance

	ServiceRepo *ServiceRepositoryImpl // Singleton
)

const (
	SUV        = VehicleType("SUV")
	HATCHBACK  = VehicleType("HATCHBACH")
	SEDAN      = VehicleType("SEDAN")
	MOTORCYCLE = VehicleType("MOTORCYCLE")
)

const (
	ON_RENT       = STATUS("ON_RENT")
	AVAILABLE     = STATUS("AVAILABLE")
	DISCARDED     = STATUS("DISCARDED")
	NOT_AVAILABLE = STATUS("NOT_AVAILABLE") // Gone for maintenance or could not rent for any other reasons
)

type Vehicle interface {
	// Would have all the getter and setter for the vehicle
	GetVehicleId() string
	GetVehicleType() VehicleType
	SetBookingId(bookingId string)
	SetStatus(status STATUS)

	Release()              // this would set the status of vehicle to available, mark current bookingId as nil, etc
	Pick(bookingId string) // this would set the booking id and status for vehicle

	Save()
}

type VehicleCommon struct {
	RegistrationNumber string // Unique Key
	EngineNumber       string // Unique Key
	RCNumber           string // Primary Key
	OwnerName          string
	OwnerAddress       string
	Type               VehicleType
	Company            string
	Brand              string
	YearOfPurchase     string
	Status             STATUS
	InternalComment    string // To justify for non availability of vehicle
	CurrentBookingId   string
}

type Car struct {
	VehicleCommon
	CurrentDriver string // driver Id
	CNGEnabled    bool
}

func (vehicle *Car) GetVehicleId() string {
	return vehicle.RCNumber
}

func (vehicle *Car) GetVehicleType() VehicleType {
	return vehicle.Type
}

func (vehicle *Car) SetStatus(status STATUS) {
	vehicle.Status = status
}

func (vehicle *Car) SetBookingId(bookingId string) {
	vehicle.CurrentBookingId = bookingId
}

func (vehicle *Car) Release() {
	vehicle.Status = AVAILABLE
	vehicle.CurrentBookingId = ""
}

func (vehicle *Car) Pick(bookingId string) {
	if bookingId == "" {
		return
	}
	vehicle.CurrentBookingId = bookingId
}

func (vehicle *Car) Save() {

	for id, _ := range VehicleData {
		if VehicleData[id].GetVehicleId() == vehicle.GetVehicleId() {
			VehicleData[id] = vehicle
		}
	}
}

type Bike struct {
	VehicleCommon
	SideStorageAvailable bool
}

func (vehicle *Bike) GetVehicleId() string {
	return vehicle.RCNumber
}

func (vehicle *Bike) GetVehicleType() VehicleType {
	return vehicle.Type
}


func (vehicle *Bike) SetStatus(status STATUS) {
	vehicle.Status = status
}

func (vehicle *Bike) SetBookingId(bookingId string) {
	vehicle.CurrentBookingId = bookingId
}

func (vehicle *Bike) Release() {
	vehicle.CurrentBookingId = ""
	vehicle.Status = AVAILABLE
}

func (vehicle *Bike) Save() {

	for id, _ := range VehicleData {
		if VehicleData[id].GetVehicleId() == vehicle.GetVehicleId() {
			VehicleData[id] = vehicle
		}
	}
}

func (vehicle *Bike) Pick(bookingId string) {
	if bookingId == "" {
		return
	}
	vehicle.CurrentBookingId = bookingId
}

// This service handles all CRUD interactions with the db for vehicle
type VehicleService interface {
	CreateVehicle() Vehicle
	GetVehicleById(id string) Vehicle
	UpdateVehicleDetails() Vehicle

	SetVehicle(vehicle Vehicle)
	GetVehicle() Vehicle
}

type VehicleServiceImpl struct {
	// Details that you may want to update
	Vehicle
}

func (vehicleService *VehicleServiceImpl) GetVehicle() Vehicle {
	return vehicleService.Vehicle
}

func (vehicleService *VehicleServiceImpl) SetVehicle(vehicle Vehicle) {
	vehicleService.Vehicle = vehicle
}

func (vehicleServiceImpl *VehicleServiceImpl) CreateVehicle() Vehicle {
	vehicleObject := vehicleServiceImpl.GetVehicle()
	VehicleData = append(VehicleData, vehicleObject)
	return vehicleObject
}

func (vehicleService *VehicleServiceImpl) GetVehicleById(id string) Vehicle {

	for _, vehicle := range VehicleData {
		if vehicle.GetVehicleId() == id {
			return vehicle
		}
	}
	return nil
}

func (vehicleService *VehicleServiceImpl) UpdateVehicleDetails() Vehicle {
	if vehicleService.GetVehicle() == nil {
		return nil
	}
	targetVehicle := vehicleService.GetVehicle()
	for idx, _ := range VehicleData {
		if VehicleData[idx].GetVehicleId() == targetVehicle.GetVehicleId() {
			VehicleData[idx] = targetVehicle
			return targetVehicle
		}
	}
	return nil
}

// All the addon services would be linked to a booking
// Service can have eligibility for each vehicle type
type ServiceRepository interface {
	GetAllServices() []Service
	GetServiceWithId(string) Service
	GetServicesForIds([]string) []Service
}

type ServiceRepositoryImpl struct {
	ServiceList []Service
}

func (serviceRepo *ServiceRepositoryImpl) GetAllServices() []Service {
	if len(serviceRepo.ServiceList) == 0 {
		//InitializeVehicleServices()
		// May return error in this case
	}
	return serviceRepo.ServiceList
}

func (serviceRepo *ServiceRepositoryImpl) GetServiceWithId(id string) Service {
	allServices := serviceRepo.GetAllServices()
	for _, service := range allServices {
		if service.GetServiceId() == id {
			return service
		}
	}
	return nil
}

func (serviceRepo *ServiceRepositoryImpl) GetServicesForIds(ids []string) []Service {
	allServices := serviceRepo.GetAllServices()
	response := []Service{}

	for _, id := range ids {
		for _, service := range allServices {
			if id == service.GetServiceId() {
				response = append(response, service)
			}
		}
	}
	return response
}

type Service interface {
	GetServiceCost() int

	// Add all the getter and setter for convenience
	GetServiceId() string
}

type AddonDriver struct {
	Id          string
	Description string
	Cost        int // Fixed cost
}

func (service *AddonDriver) GetServiceCost() int {
	return service.Cost
}

func (service *AddonDriver) GetServiceId() string {
	return service.Id
}

type FueledTank struct {
	Id          string
	Description string
	Cost        int
}

func (service *FueledTank) GetServiceCost() int {
	return service.Cost
}

func (service *FueledTank) GetServiceId() string {
	return service.Id
}

type Insurance struct {
	Id          string
	Description string
	Cost        int
}

func (service *Insurance) GetServiceCost() int {
	return service.Cost
}

func (service *Insurance) GetServiceId() string {
	return service.Id
}

type Inventory struct {
	Id              string // Primary Key
	VehicleType     VehicleType
	PricePerVehicle int
	TotalCount      int
	Available       int
	OnRentDuty      int
	Others          int
}

func InitializeVehicleServices() {
	AddOnDriverService = &AddonDriver{
		Id:          "1",
		Description: "Get driver with the booking",
		Cost:        1000,
	}

	FueledTankService = &FueledTank{
		Id:          "2",
		Description: "Get vehicle with full tank",
		Cost:        500,
	}

	InsuranceService = &Insurance{
		Id:          "3",
		Description: "Insurance for the vehicle",
		Cost:        300,
	}

	ServiceRepo = &ServiceRepositoryImpl{
		[]Service{
			InsuranceService,
			FueledTankService,
			AddOnDriverService,
		},
	}
}

// Functionality Required
// - Return a Vehicle
//		- This would mark the vehicle status as available and update the inventory accordingly.
// - Provide a Vehicle to a person. Input: BookingId, userId, VehicleObject.
//		- Booking details would be updated with exact vehicle handed over for the booking
// 		- Vehicle would be updated with the booking details
// 		- Inventory Table would be accordingly updated using the vehicle status.
