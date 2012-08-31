package main

import "fmt"

var console = Console("")

type Console string

func (self Console) log(s interface{}) {
	fmt.Println(s)
}


