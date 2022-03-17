package model

import (
	"errors"
)

type Machine interface {
	Start() // Initialize all slot types, Update the inventory it has,
	Stop()

	GetCurrentOrder() Order
	CreateNewOrder(int, int) (int, error) // use quantity and slot number, returns amount and error if any
	ValidateOrder() error
	DumpCurrentOrderToDB()
	AcceptPaymentForOrder(details PaymentDetails) error

	Dispense() error // Would release the items corresponding to the current order
	AddItemToSlot(item Item, quantity int, slotNumber int) error
}
type VendingMachine struct {
	CurrentOrder *Order
	Slots        []Slot
	SlotMap      map[int]Slot
	State        string // ON/OFF/Unavailable
}

func (machine *VendingMachine) Start() {
	slotMap := make(map[int]Slot)
	machine.CurrentOrder = nil
	machine.Slots = InitializeAllSlots()
	for _, slot := range machine.Slots {
		slotMap[slot.Number] = slot
	}
	machine.SlotMap = slotMap
	machine.State = "ACTIVE"
}

func (machine *VendingMachine) Stop() {
	if machine.CurrentOrder != nil {
		machine.DumpCurrentOrderToDB(true)
	}
	machine.CurrentOrder = nil
	machine.State = "OFF"
}

func (machine *VendingMachine) GetCurrentOrder() Order {
	return *machine.CurrentOrder
}

func (machine *VendingMachine) CreateNewOrder(slotNumber int, quantity int) (int, error) {
	order := Order{Id: "3", SlotNumber: slotNumber, Quantity: quantity}
	machine.CurrentOrder = &order
	err := machine.ValidateOrder()
	if err != nil {
		machine.CurrentOrder = nil
		return 0, err
	}
	orderItem := machine.SlotMap[slotNumber].Item
	paymentAmount := (&orderItem).GetPrice() * quantity
	machine.CurrentOrder.PaymentAmount = paymentAmount
	return paymentAmount, nil
}

func (machine *VendingMachine) ValidateOrder() error {
	if machine.CurrentOrder.Quantity == 0 || machine.CurrentOrder.SlotNumber == 0 {
		return errors.New("Missing Order Details")
	}
	if slot, ok := machine.SlotMap[machine.CurrentOrder.SlotNumber]; ok {
		if slot.Quantity < machine.CurrentOrder.Quantity {
			return errors.New("Order quantity exceeds the available quantity")
		}
	} else {
		return errors.New("Invalid Slot Number provided for order")
	}

	return nil
}

func (machine *VendingMachine) DumpCurrentOrderToDB(markCurrenOrderNil bool) {
	exists := false
	for id, order := range OrdersDB {
		if order.Id == machine.CurrentOrder.Id {
			OrdersDB[id] = *machine.CurrentOrder
			exists = true
		}
	}
	if !exists {
		OrdersDB = append(OrdersDB, *machine.CurrentOrder)
	}
	if markCurrenOrderNil {
		machine.CurrentOrder = nil
	}
}
func (machine *VendingMachine) AcceptPaymentForOrder(details PaymentDetails) error {
	if machine.CurrentOrder == nil {
		return errors.New("No existing order!")
	}
	machine.CurrentOrder.PaymentDetails = details
	return nil
}

func (machine *VendingMachine) Dispense() error {
	order := machine.CurrentOrder

	if order == nil {
		return errors.New("Error in dispensing! Unable to find order!")
	}

	if order.SlotNumber != 0 && order.Quantity != 0 {
		// START MUTEX : LOCKED
		slot := machine.SlotMap[order.SlotNumber]
		if slot.Quantity < order.Quantity {
			return errors.New("Order quantity exceeds the available quantity")
		}
		slot.Quantity = slot.Quantity - order.Quantity
		machine.SlotMap[order.SlotNumber] = slot
		for id, existingslot := range machine.Slots {
			if existingslot.Number == order.SlotNumber {
				machine.Slots[id] = slot
			}
		}
		// END MUTEX : UNLOCKED
	} else {
		return errors.New("Invalid order details")
	}
	machine.DumpCurrentOrderToDB(true)
	machine.CurrentOrder = nil

	return nil
}

func (machine *VendingMachine) AddItemToSlot(item Item, quantity int, slotNumber int) error {
	machine.Stop()
	// Check if there is space in slot
	slot := machine.SlotMap[slotNumber]
	if slot.Item.ItemId != item.ItemId {
		return errors.New("Unmatched Item")
	}
	slot.Quantity = min(slot.Limit, slot.Quantity+quantity)
	machine.SlotMap[slotNumber] = slot
	for id, existingSlot := range machine.Slots {
		if slot.Number == existingSlot.Number {
			machine.Slots[id] = slot
		}
	}
	return nil
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Can use getter and setter for this. No need for an interface in this case.
type Order struct {
	Id             string
	SlotNumber     int
	Quantity       int
	PaymentAmount  int
	PaymentDetails PaymentDetails
}

type PaymentDetails struct {
	PaymentId     string
	PaymentAmount int
	Method        string // could be an enum
}
