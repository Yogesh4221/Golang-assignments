package main

import "fmt"

type foo interface {
	bar() string
}

type newfoo interface {
	baz() string
}

// We want a `foo` interface type, but if that valid type can also do the new
// behaviour, then we'll execute that behaviour...

func doThing(f foo) {
	if nf, ok := f.(newfoo); ok {
		fmt.Println("baz:", nf.baz())
	}
	fmt.Println("bar:", f.bar())
}

// Original concrete implementation...

type point struct {
	X, Y int
}

func (p point) bar() string {
	return fmt.Sprintf("p=%d, y=%d", p.X, p.Y)
}

// New concrete implementation of `point` struct (has the new method)...

type newpoint struct {
	point
}

func (np newpoint) baz() string {
	return fmt.Sprintf("np !!! %d, ny !!! %d", np.X, np.Y)
}

func main() {
	var pt point
	pt.X = 1
	pt.Y = 2

	doThing(pt)

	var npt newpoint
	npt.X = 3
	npt.Y = 4

	doThing(npt)
}
