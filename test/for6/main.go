package main

import "log"


type Address struct {
	Number int
	Street string
	Name string
	ZipCode int
}

func main() {	
	adds := []Address{}

	for i:=0; i<10; i++ {
		adds = append(adds, Address {21, "Jump st", "John Doe", 10001})
	}
}