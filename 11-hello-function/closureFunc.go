package main

func closureFunc() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
