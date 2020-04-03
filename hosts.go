package main

import (
	"fmt"

	"github.com/HewlettPackard/simplivity-go/ovc"
)

func main() {
	/*
		var (
			hostName   = "OmniStackVC-10-3-4-72"
			hostByName *ovc.Host
		)
	*/

	//Create an ovc client
	client, err := ovc.NewClient("usradmin@v0004.local", "HP!nvent123", "10.3.4.72", "")
	if err != nil {
		fmt.Println(err)
	}

	//Get all hosts resources without Filter
	hostList, err := client.Hosts.GetAll(ovc.GetAllParams{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Get all Hosts without params ")
	fmt.Println("----------------------------------------")
	for _, hosts := range hostList.Members {
		fmt.Println(hosts.Name + "\n")
	}
}
