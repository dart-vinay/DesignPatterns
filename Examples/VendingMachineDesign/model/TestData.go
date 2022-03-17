package model

var OrdersDB = []Order{}
var Inventory = []Slot{
	{
		Number: 1,
		Type:   CanT1,
		Item: Item{
			"1",
			"Coffee",
			35,
			"",
		},
		Limit:    10,
		Quantity: 10,
	}, {
		Number: 2,
		Type:   CanT1,
		Item: Item{
			"2",
			"Cola",
			25,
			"",
		},
		Limit:    10,
		Quantity: 10,
	},
	{
		Number: 3,
		Type:   Bottle25,
		Item: Item{
			"3",
			"ThumsUp 600",
			45,
			"",
		},
		Limit:    8,
		Quantity: 6,
	},
}
