package main

type X struct{ x int }

func (x *X) F1() { println("(*X).F1:", x.x) }
func (x *X) F2() { println("(*X).F2") }

type I interface {
	F1()
	F2()
}

func main() {
	var x interface{}
	x = int32(123456)

	// Let's try an interface-to-value assertion.
	if x, ok := x.(int32); ok {
		println("i2v:", x)
	}
	if x, ok := x.(int64); ok {
		println("i2v:", x)
	}

	// This will fail the assertion.
	if i, ok := x.(I); ok {
		i.F1()
		_ = i
	} else {
		println("!")
	}

	// Assign an *X, which should pass the assertion.
	x_ := new(X)
	x_.x = 123456
	x = x_ //&X{x: 123456}
	if i, ok := x.(I); ok {
		i.F1()
		_ = i
	} else {
		println("!")
	}
}
