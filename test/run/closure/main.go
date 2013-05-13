package main


func main() {
	x := 0
	f := func() {
		x = x + 1
	}
	f()
	console.log(x)
}






