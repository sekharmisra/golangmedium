package main

import "fmt"

func main() {

	//gujarati Om taken from https://www.ssec.wisc.edu/~tomw/java/unicode.html#x0A80
	om := 'ૐ'
	fmt.Printf("Character: %c, Value: %v, Unicode: %U, Type: %T", om, om, om, om)
}
