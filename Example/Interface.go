package main

import "fmt"

type Cloud interface {
	VMProvisioining(string) bool
	VMTerminate(string) bool
}

type VM struct {
	Name    string
	State   string
	Running bool
}

func (vm *VM) VMProvisioining(name string) bool {
	vm.Name = name
	if vm.Running == true {
		return true
	} else {
		return false
	}

}

func (vm *VM) VMTerminate(name string) bool {
	vm.Name = name
	if vm.Running == true {
		vm.Running = false
		return vm.Running
	}
	return vm.Running
}

func main() {

	MasterVm := VM{
		Name:    "K8s-Master",
		State:   "Running",
		Running: false,
	}

	state := MasterVm.VMProvisioining("K8S-MANAGER")
	if state {
		fmt.Println(MasterVm.Name, "Vm is running")
	} else {
		fmt.Println(MasterVm.Name, "Vm is Not running")
	}

	state = MasterVm.VMTerminate(MasterVm.Name)

	if !state {
		fmt.Println(MasterVm.Name, "Vm is Not running")
	}
}
