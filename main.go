package main

import "fmt"

type Color int

// fmt.Stringer
func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "Blue"
	case ColorBlack:
		return "Black"
	case ColorYellow:
		return "Yellow"
	case ColorPink:
		return "Pink"
	default:
		panic("invalid color given")
	}
}

const (
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func main() {
	fmt.Println("the color is ", ColorBlack)
}
