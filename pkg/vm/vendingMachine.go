package vm

import "fmt"

type Item struct {
	Quantity int
	Price    int
	Name     string
}

type VendingMachine struct {
	Items   map[string]*Item
	State   State
	Balance int
}

func NewVendingMachine() *VendingMachine {

	return &VendingMachine{
		Items:   make(map[string]*Item),
		Balance: 0,
		State:   &IdleState{},
	}
}

func (vm *VendingMachine) AddItem(item Item) {

	if vm.Items == nil {
		vm.Items = make(map[string]*Item)
	}
	vm.Items[item.Name] = &item
	fmt.Println("Added item:", item.Name, "Quantity:", item.Quantity, "Price:", item.Price)
}

func (vm *VendingMachine) SetState(s State) {
	vm.State = s
}

func (vm *VendingMachine) SelectItem(itemName string) error {
	return vm.State.SelectItem(itemName, vm)

}
func (vm *VendingMachine) InsertMoney(money int) error {
	return vm.State.InsertMoney(money, vm)

}

func (vm *VendingMachine) Dispense() error {
	return vm.State.DispenseItem(vm)
}
