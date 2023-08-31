// packages are collection of files arranged in single package. You need to init go module before compiling the program!
package main

import (
	"fmt"

	"github.com/google/uuid"
)

func hellouuid() {
	fmt.Println(uuid.New().String())

}
