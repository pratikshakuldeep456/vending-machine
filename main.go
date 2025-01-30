package main

import (
	//"https://pkg.go.dev/pratikshakuldeep456/vending-machine/pkg/vm#Item"
	"fmt"
	"pratikshakuldeep456/vending-machine/pkg/vm"
)

func main() {

	vendingMachine := vm.NewVendingMachine()

	// Add items
	vendingMachine.AddItem(vm.Item{Quantity: 10, Price: 10, Name: "uncleChips"})
	vendingMachine.AddItem(vm.Item{Quantity: 10, Price: 15, Name: "kitkat"})

	if err := vendingMachine.SelectItem("uncleChips"); err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := vendingMachine.InsertMoney(10); err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := vendingMachine.Dispense(); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
