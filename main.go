package main

func main() {
	lh := learnHttp{}
	li := learnInterface{}
	ls := learnStruct{}
	lm := learnMap{}
	lc := learnConcurrency{}
	lh.executeMain()
	li.executeMain()
	ls.executeMain()
	lm.executeMain()
	lc.executeMain()
}

// t := triangle{base: 1, height: 2}
// s := square{sideLength: 2}
// printArea(t)
// printArea(s)

// type triangle struct {
// 	base   float64
// 	height float64
// }

// type square struct {
// 	sideLength float64
// }

// type shape interface {
// 	getArea() float64
// }

// func (t triangle) getArea() float64 {
// 	return t.base * t.height * 0.5
// }

// func (s square) getArea() float64 {
// 	return s.sideLength * s.sideLength
// }

// func printArea(s shape) {
// 	fmt.Println("Area of this share is ", s.getArea())
// }
