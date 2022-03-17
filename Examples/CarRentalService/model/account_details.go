package model

import (
	"errors"
	"time"
)

// This is the way multiple account creation should be handled.
// This is in a way an abstract factory pattern where appropriate handlers are fetched who in turn are responsible for creation and other relevant tasks for that type of account.

type AccountType string

const (
	USER   = AccountType("USER")
	DRIVER = AccountType("DRIVER")
	ADMIN  = AccountType("ADMIN")
)

// All the account related activity would be carried through this interface
type AccountHandler interface {
	CreateAccount(account Account) (Account, error)
	DeleteAccount(accountId string)
	GetAccount(accountId string) (Account, error)

	// Additional Functionalities
	// UpdateAccountDetails(Account) (Account, error)
	//
}

// This is the common account interface. For attributes not common to all the accounts we would separately implement the functionality of getters and setters.
type Account interface {
	GetAccountType() AccountType
	// Add all other relevant getter and setter methods to provide complete functionality.
}

type AccountCommonInfo struct {
	Id           string
	Email        string
	PhoneNumber  string
	CreatedAt    time.Time
	CreatedBy    string
	Active       bool
	LastAccessed time.Time
	AccountType  AccountType
}

type LicenceInfo struct {
	LicenceNumber string
	IssuedOn      time.Time
	Expiry        time.Time
	LicenceType   string
}

type UserAccount struct {
	AccountCommonInfo
	LicenceInfo
}

type AdminAccount struct {
	AccountCommonInfo
}

type DriverAccount struct {
	AccountCommonInfo
	LicenceInfo
	HiredOn   time.Time
	Available bool
}

func (account *UserAccount) GetAccountType() AccountType {
	return account.AccountType
}

func (account *UserAccount) GetLicenseInfo() LicenceInfo {
	return account.LicenceInfo
}

func (account *AdminAccount) GetAccountType() AccountType {
	return account.AccountType
}

func (account *DriverAccount) GetAccountType() AccountType {
	return account.AccountType
}

func (account *DriverAccount) GetLicenseInfo() LicenceInfo {
	return account.LicenceInfo
}

type UserAccountHandler struct {
}

type DriverAccountHandler struct {
}

type AdminAccountHandler struct {
}

func (userAccountHandler *UserAccountHandler) CreateAccount(userAccount Account) (Account, error) {
	if userAccount.GetAccountType() != USER {
		return userAccount, errors.New("Account type incorrect")
	}

	UserData = append(UserData, *(userAccount.(*UserAccount)))
	return userAccount, nil
}

func (userAccount *UserAccountHandler) DeleteAccount(accountId string) {
	modifiedAccountList := []UserAccount{}
	for _, account := range UserData {
		if account.Id != accountId {
			modifiedAccountList = append(modifiedAccountList, account)
		}
	}
	UserData = modifiedAccountList
}

func (userAccountHandler *UserAccountHandler) GetAccount(accountId string) (Account, error) {
	account := UserAccount{}
	for _, a := range UserData {
		if a.Id == accountId {
			account = a
		}
	}
	return &account, nil
}

func (adminAccountHandler *AdminAccountHandler) CreateAccount(adminAccount Account) (Account, error) {
	if adminAccount.GetAccountType() != ADMIN {
		return adminAccount, errors.New("Account type incorrect")
	}
	AdminData = append(AdminData, *(adminAccount.(*AdminAccount)))
	return adminAccount, nil
}

func (adminAccountHandler *AdminAccountHandler) DeleteAccount(accountId string) {
	modifiedAccountList := []AdminAccount{}

	for _, account := range AdminData {
		if account.Id != accountId {
			modifiedAccountList = append(modifiedAccountList, account)
		}
	}

	AdminData = modifiedAccountList
}

func (adminAccountHandler *AdminAccountHandler) GetAccount(accountId string) (Account, error) {
	account := AdminAccount{}
	for _, a := range AdminData {
		if a.Id == accountId {
			account = a
		}
	}
	return &account, nil
}

func (driverAccountHandler *DriverAccountHandler) CreateAccount(driverAccount Account) (Account, error) {
	if driverAccount.GetAccountType() != DRIVER {
		return driverAccount, errors.New("Account type incorrect")
	}
	DriverData = append(DriverData, *(driverAccount.(*DriverAccount)))

	return driverAccount, nil
}

func (driverAccountHandler *DriverAccountHandler) DeleteAccount(accountId string) {
	modifiedAccountList := []DriverAccount{}

	for _, account := range DriverData {
		if account.Id != accountId {
			modifiedAccountList = append(modifiedAccountList, account)
		}
	}

	DriverData = modifiedAccountList
}

func (driverAccountHandler *DriverAccountHandler) GetAccount(accountId string) (Account, error) {
	account := DriverAccount{}
	for _, a := range DriverData {
		if a.Id == accountId {
			account = a
		}
	}
	return &account, nil
}

func GetAccountHandler(accountType AccountType) AccountHandler {
	switch accountType {
	case USER:
		return new(UserAccountHandler)
	case ADMIN:
		return new(AdminAccountHandler)
	case DRIVER:
		return new(DriverAccountHandler)
	default:
		return new(UserAccountHandler)
	}
}

func CreateAccount(account Account, accountType AccountType) (Account, error) {
	accountHandler := GetAccountHandler(accountType)
	account, err := accountHandler.CreateAccount(account)
	return account, err
}

// Functionality Required
// - All bookings for the user in the searched time frame

type FetchUserBookings struct {
	UserId string
	Status STATUS
}

func (req *FetchUserBookings) FetchBookingsForUser() ([]Booking, error) {

	// Get booking details for user from db and return filtered results.
	return []Booking{}, nil
}
