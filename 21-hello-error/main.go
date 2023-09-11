package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.ReadFile("myfile.txt")

	if err != nil {
		//print an error
		fmt.Println(err)

		//log an error - prints the error to standard logger
		log.Println(err)

		//log fatal
		log.Fatalln(err)

		panic(err)
	}
}
