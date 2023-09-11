package main

import "strings"

func addVariadicFunc(elements []string) string {
	return strings.Join(elements, "-")
}

func addSlice() []string {
	nameSlice := []string{}
	nameSlice = append(nameSlice, "Sekhar")
	nameSlice = append(nameSlice, "Misra")
	nameSlice = append(nameSlice, "Aarna")
	nameSlice = append(nameSlice, "Misra")
	nameSlice = append(nameSlice, "Kunjan")
	nameSlice = append(nameSlice, "Misra")

	return nameSlice
}
