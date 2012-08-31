package main

func ifsquare(x, y int) int {
	if x < 0 {
		return x * y
	} else {
		return y * x
	}
	return 0
}

func main() {	
	ifsquare(3, 4)
}