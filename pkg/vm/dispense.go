package vm

import (
	"errors"
	"fmt"
)

type DispenseState struct {
	Item Item
}

func (ds *DispenseState) SelectItem(itemName string, vm *VendingMachine) error {

	return errors.New("item already selected ")

}

func (ds *DispenseState) InsertMoney(amount int, vm *VendingMachine) error {

	return errors.New("payment already done")
}

func (ds *DispenseState) DispenseItem(vm *VendingMachine) error {

	key, val := vm.Items[ds.Item.Name]
	if !val || key.Quantity <= 0 {
		return errors.New("item is out of stock")
	}
	key.Quantity--

	fmt.Printf("Dispensed: %s. Remaining stock: %d\n", ds.Item.Name, key.Quantity)
	vm.SetState(&IdleState{})

	return nil
}
