package main

import "log"

var console = Console("")

type Console string

func (self Console) log(s interface{}) {
	log.Println(s)
}


