package vm

import (
	"errors"
	"fmt"
)

type PaymentState struct {
	Item Item
}

func (ps *PaymentState) SelectItem(itemName string, vm *VendingMachine) error {
	return errors.New("item already selected ")
}

func (ps *PaymentState) InsertMoney(amount int, vm *VendingMachine) error {

	inventory := vm.Items

	if amount < inventory[ps.Item.Name].Price {
		return errors.New("insufficient money inserted")
	}
	vm.Balance += amount
	fmt.Printf("Money inserted: %d. Dispensing %s...\n", amount, ps.Item.Name)

	vm.SetState(&DispenseState{Item: ps.Item})
	return nil

}

func (ps *PaymentState) DispenseItem(vm *VendingMachine) error {
	return errors.New("payment not done")

}
