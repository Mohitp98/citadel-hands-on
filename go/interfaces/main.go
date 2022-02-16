package main

import "fmt"

// interface acts as a blueprint for "Methods Sets"
// It holds only the signature of methods
type shape interface {
	area() float64
	perimeter() float64
}

type Square struct {
	side float64
}

// Method sets for Square struct
func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return 4 * s.side
}

type Reactangle struct {
	length, breadth float64
}

// Method sets for Square struct
// this is normal receiver function that
// follows signature mentioned in the shape interface
func (r Reactangle) area() float64 {
	return r.breadth * r.length
}

func (r Reactangle) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

// this is where interfaces shines!
func printData(s shape) {
	fmt.Printf("Interface: %T | Values: %v\n", s, s)
	fmt.Println("Area :", s.area())
	fmt.Println("Perimeter :", s.perimeter())
}

func main() {
	s := Square{
		side: 45.6,
	}
	printData(s)

	r := Reactangle{
		length:  10,
		breadth: 3,
	}
	printData(r)

}
