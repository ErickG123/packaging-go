package main

import (
	"fmt"
	"packaging/math"
)

func main() {
	m := math.NewMath(1, 2)
	// m := math.Math{}
	fmt.Println(m.Add())
}
