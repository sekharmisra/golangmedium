package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name    string `json:"user_name"`
	Age     int    `json:"user_age"`
	Address string `json:"user_address"`
	ZipCode string `json:"user_postalcode"`
}

func marshalFunc() {
	//Marshal
	user := User{
		Name:    "Sekhar Misra",
		Age:     40,
		Address: "Waterloo",
		ZipCode: "7",
	}
	result, err := json.Marshal(user)

	if err != nil {
		log.Fatal(err)
	} else {
		//print the result
		fmt.Println(string(result))
	}

	//Unmarshal
	var umUser User
	err = json.Unmarshal(result, &umUser)
	if err != nil {
		log.Fatal(err)
	} else {
		//print the result
		fmt.Println(umUser)
	}
}
