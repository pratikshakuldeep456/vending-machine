package vm

type State interface {
	SelectItem(itemName string, vm *VendingMachine) error
	InsertMoney(amount int, vm *VendingMachine) error
	DispenseItem(vm *VendingMachine) error
}
