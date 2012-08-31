package main

func ifsquare(x, y int) int {
	if x < 0 {
		return x * y
	}
	return y * x
}

func main() {	
	ifsquare(3, 4)
}