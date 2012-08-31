package main

func count(a, b int) int {
	for i:=a; i<b; i++ {
		a = i
	}
	return a
}

func main() {	
	t := count(0,10)
	console.log(t)
}