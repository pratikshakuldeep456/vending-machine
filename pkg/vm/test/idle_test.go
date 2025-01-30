package test

import (
	"fmt"
	"pratikshakuldeep456/vending-machine/pkg/vm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newVendingMachineIdle() *vm.VendingMachine {
	vendingMachine := vm.NewVendingMachine()

	item := vm.Item{
		Quantity: 1,
		Price:    10,
		Name:     "Coke",
	}
	vendingMachine.AddItem(item)
	return vendingMachine
}

func TestSelectItemIdleState(t *testing.T) {
	vm := newVendingMachineIdle()

	err := vm.State.SelectItem("Coke", vm)

	assert.NoError(t, err, "expected no error")
}

func TestSelectItemIdleStateWithWrongItem(t *testing.T) {
	vm := newVendingMachineIdle()

	err := vm.State.SelectItem("Slice", vm)

	assert.EqualError(t, err, fmt.Sprintf("item %s is not available", "Slice"))
}

func TestInsertMoneyIdleState(t *testing.T) {
	vm := newVendingMachineIdle()
	err := vm.State.InsertMoney(10, vm)

	assert.EqualError(t, err, "item not selected ")

}

func TestDispenseItemIdleState(t *testing.T) {
	vm := newVendingMachineIdle()
	err := vm.State.DispenseItem(vm)
	assert.EqualError(t, err, "item not selected ")
}
