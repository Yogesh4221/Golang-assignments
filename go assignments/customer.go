package main

import "fmt"

func main() {
	var customer = map[int]string{
		1001: "Vijay",
		1002: "Ajith",
		1003: "surya",
	}

	printCustomer(customer, 1001)
	printCustomer(customer, 1010)

	if iscustomerExists(customer, 1002) {
		fmt.Println("customer Id 1002 found")
	}
}

func printCustomer(customer map[int]string, customerId int) {
	if name, ok := customer[customerId]; ok {
		fmt.Printf("name = %s, ok = %v\n", name, ok)
	} else {
		fmt.Printf("customer %d not found\n", customerId)
	}
}

func iscustomerExists(customer map[int]string, customerId int) bool {
	_, ok := customer[customerId]
	return ok
}
