package test

import (
	"pratikshakuldeep456/vending-machine/pkg/vm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectItemDispenseState(t *testing.T) {
	vm := newVendingMachineIdle()

	err := vm.State.SelectItem("Coke", vm)
	err = vm.State.InsertMoney(10, vm)

	err = vm.State.SelectItem("Coke", vm)

	assert.EqualError(t, err, "item already selected ")

}

func TestInsertMoneyDispenseState(t *testing.T) {
	vm := newVendingMachineIdle()

	err := vm.State.SelectItem("Coke", vm)
	err = vm.State.InsertMoney(10, vm)
	err = vm.State.InsertMoney(10, vm)

	assert.EqualError(t, err, "payment already done")
}

func TestDispenseItemDispenseState(t *testing.T) {
	vm := newVendingMachineIdle()

	err := vm.State.SelectItem("Coke", vm)
	err = vm.State.InsertMoney(10, vm)
	err = vm.State.DispenseItem(vm)

	assert.NoError(t, err)

}

func TestDispenseItemDispenseStateNoQuantity(t *testing.T) {
	vM := vm.NewVendingMachine()

	item := vm.Item{
		Quantity: 0,
		Price:    5,
		Name:     "Coke",
	}

	vM.AddItem(item)
	err := vM.State.SelectItem("Coke", vM)
	err = vM.State.InsertMoney(10, vM)
	err = vM.State.DispenseItem(vM)

	assert.EqualError(t, err, "item is out of stock")

}
