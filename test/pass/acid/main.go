package main

/* acidtest1 */
// comment Group
// import "fmt"

type Num int

var c = 1

// const (
// 	a Num = iota
// 	b
// )

type Eq interface {
	Same(Eq) bool
}

type point struct {
	Num
	x, y float64
	z    float64 "up/down"
}

func (p *point) X() float64 {
	return -(*p).x
}

func all(...bool) {}

func main() {
	type Point point
	block := make(chan bool)
	defer fmt.Println("Done")
	item := []interface{}{`abc`}[0]
	switch item.(type) {
	case string:
		s := item.(string)
		if s[0:1] == ("a") {
			for i, c := range s {
				for j := range s {
					for j > 0 {
						for k := 0; k < j; k++ {
							j--
							fmt.Println(i, j, c)
						}
					}
				}
			}
		} else {
			switch {
			case true:
				_ = map[string]string{"key": "val", "a": "b"}
			default:
				p := &point{}
				p.X()
			}
			break
		}
	}
	go func() { block <- true }()
	select { case <-block: goto finished }
finished:
}




