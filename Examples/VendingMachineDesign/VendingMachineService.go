package VendingMachineDesign

import (
	"DesignPatterns/Examples/VendingMachineDesign/model"
	"github.com/labstack/gommon/log"
)

func VendingMachineService() {

	machine := model.VendingMachine{}
	machine.Start()

	amount, err := machine.CreateNewOrder(1, 5)
	if err != nil {
		log.Errorf("Error creating order %v", err)
	}
	log.Info(amount)
	machine.DumpCurrentOrderToDB(false)
	orders := model.OrdersDB
	err = machine.AcceptPaymentForOrder(model.PaymentDetails{"55322", machine.CurrentOrder.PaymentAmount, ""})
	if err != nil {
		log.Info(err)
	}
	machine.DumpCurrentOrderToDB(false)
	machine.Dispense()
	machine.AddItemToSlot(model.Item{"1","Dummy", 35,""},2,1)
	machine.AddItemToSlot(model.Item{"1","Dummy", 35,""},50,1)
	log.Info(orders)
}
