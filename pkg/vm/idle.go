package vm

import (
	"errors"
	"fmt"
)

type IdleState struct {
}

func (is *IdleState) SelectItem(itemName string, vm *VendingMachine) error {
	fmt.Println("Items in vending machine:", vm.Items)

	inventory := vm.Items

	i, exists := inventory[itemName]
	if !exists {

		return fmt.Errorf("item %s is not available", itemName)
	}

	vm.SetState(&PaymentState{Item: *i})
	fmt.Printf("Item %s selected. Please insert %d units of money.\n", itemName, i.Price)

	return nil

}

func (is *IdleState) InsertMoney(amount int, vm *VendingMachine) error {

	return errors.New("item not selected ")

}

func (is *IdleState) DispenseItem(vm *VendingMachine) error {
	return errors.New("item not selected ")
}
