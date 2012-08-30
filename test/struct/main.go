package main

type Point struct {
	x int
	y int
}

func main() {	
	p := Point{1, 2}
	log.Println(p)
}