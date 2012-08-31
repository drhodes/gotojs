package main

import "log"

func main() {	
	for a, b := range []int{1,2,3,4} {
		log.Println(a + b)
	}
}