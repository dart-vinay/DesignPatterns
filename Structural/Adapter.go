package Structural

import (
	"fmt"
)

// It allows incompatible interfaces to collaborate by implementing both old and new interfaces.
// Example Legacy Printer wrapped in a NewAgePrinter which can store the data. Similarly there can be multiple examples like:
// - Data converter adapter. Use when shifting from one data format to another.
// - At places where we need compatibility of old interface with the newer ones.
// Also known as the WRAPPER

type LegacyPrinter interface {
	Print(string)
}

type NewAgePrinter interface {
	PrintStored()
}

type OldPrinter struct {
}

type NewPrinter struct {
	Data string
}

func (printer *OldPrinter) Print(data string) {
	fmt.Print(data)
}

func (printer *NewPrinter) PrintStored() {
	fmt.Print(printer.Data)
}

// It implements both new and old interfaces.
type PrinterAdapter struct {
	*OldPrinter
	Data string
}

// A new printer would have data stored in the object itself. So to print something from old printer wrap it in the printer adapter. This way there would be only one object exposed to the client.
func (printerAdapter *PrinterAdapter) PrintStored() {
	if printerAdapter.OldPrinter == nil {
		fmt.Print(printerAdapter.Data)
	} else {
		printerAdapter.OldPrinter.Print("Adapted: " + printerAdapter.Data)
	}
}
