package main

func main() {
	console.log("1")
	defer console.log("2")
	defer console.log("3")
}
