package main

import "counts"

func main() {
	m := counts.NewMath(1, 2)
	println(m.Add())
}
