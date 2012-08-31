package main

type Point struct {
	x, y float32
}

func (this Point) Add (other Point) Point {
	this.x += other.x
	this.y += other.y
	return this
}

func (this Point) Equal (other Point) bool {
	return this.x == other.x && this.y == other.y
}

func main() {	
	p1 := Point{1, 0}
	p2 := Point{0, 1}
	p3 := Point{1, 1}
	console.log(p1.Add(p2).Equal(p3))
}