package main

import (
	"fmt"
	"github.com/HewlettPackard/simplivity-go/ovc"
)

func main() {
	var (
		userName    = "usradmin@v0004.local"
		password    = "HP!nvent123"
		hostAddress = "10.3.4.72"
	)

	client, err := ovc.NewClient(userName, password, hostAddress, "")

	if err != nil {
		fmt.Println(err)
	}

	vmList, err := client.VirtualMachines.GetAll(ovc.GetAllParams{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Get list of all VMs")
	fmt.Println("----------------------------------------")
	for _, vms := range vmList.Members {
		fmt.Println(vms.Name + "\n")
	}
}
