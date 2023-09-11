package main

func main() {
	sq := square{
		side: 10,
	}

	cir := circle{
		radius: 10.1,
	}

	printShape(sq)
	printShape(cir)

	add := Addition{
		a: 10,
		b: 20,
	}

	mul := Multiplication{
		c: 10,
		d: 20,
	}

	printNumbers(add)
	printNumbers(mul)

}
