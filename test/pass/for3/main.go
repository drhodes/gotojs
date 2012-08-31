package main

import "log"

func main() {	
	temp := []int{1,2,3,4}
	for a, b := range temp {
		log.Println(a + b)
	}
}