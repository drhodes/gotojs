package main

func square(x int) int {
	return x * x
}

func main() {	
	defer square(4)	
	temp := square(4)
	temp = square(temp)
}






