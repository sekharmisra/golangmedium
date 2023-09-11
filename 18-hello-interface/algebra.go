package main

import "fmt"

type Addition struct {
	a, b float64
}

type Multiplication struct {
	c, d float64
}

func (addition Addition) calculate() float64 {
	return addition.a + addition.b
}

func (multiplication Multiplication) calculate() float64 {
	return multiplication.c * multiplication.d
}

type Algebra interface {
	calculate() float64
}

func printNumbers(algebra Algebra) {
	fmt.Println(algebra.calculate())
}
