package main

type Address struct {
	Number int
	Street string
	Name string
	ZipCode int
}

func main() {	
	adds := []Address{}

	for i:=0; i<10; i++ {
		temp := Address {21, "Jump st", "John Doe", 10001}
		adds = append(adds, temp)
		console.log(adds)
	}
}