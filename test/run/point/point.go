package main

type point struct {
	x, y int
}

func (p point) add(other point) point {
	return point{p.x + other.x, p.y + other.y}
}

func main() {
	p1 := point{3,5}
	console.log(p1.x)
	console.log(p1.y)

	p2 := point{7,11}
	console.log(p2.x)
	console.log(p2.y)
	
	p3 := p1.add(p2)
	console.log(p3.x)
	console.log(p3.y)

	count := 0
	for count < 10 {
		count++
		p3 = p3.add(p1)
		console.log(p3.x)
	}
}






