package main

import "fmt"

func addStringFunc(firstname string, lastname string) string {
	return fmt.Sprint(firstname, " ", lastname)
}
