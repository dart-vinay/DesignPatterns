package model

// So there could be multiple Item types in our machine: For example:
// Type of Slots :
// Slots for Packaged solid foods
// Slots for drink cans
//

type SlotType struct {
	Name          string
	Height        int
	Width         int
	WeightInGrams int
}

var (
	CanT1    = SlotType{"CanT1", 15, 5, 300}
	Bottle25 = SlotType{"BottleT1", 20, 8, 600}
)

type Item struct {
	ItemId   string
	ItemName string
	Price    int
	Type     string // Would go with a particular slot type only
}

func (item *Item) GetPrice() int {
	return item.Price
}

type Slot struct {
	Number   int
	Type     SlotType
	Item     Item
	Limit    int
	Quantity int
}

func InitializeAllSlots() []Slot {
	return Inventory
}

// Press button. Make Payment. Dispense.
