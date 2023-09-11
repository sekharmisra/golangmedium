package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name    string `json:"user_name"`
	Age     int    `json:"user_age"`
	Address string `json:"user_address"`
	ZipCode string `json:"user_postalcode"`
}

func encoderDecoderFunc() {
	user := User{
		Name:    "Sekhar Misra",
		Age:     40,
		Address: "Waterloo",
		ZipCode: "7",
	}

	//Encoding
	json.NewEncoder(os.Stdout).Encode(user)

	//Decoding
	var dUser User
	rdr := strings.NewReader(`{"user_name":"Kunjan Misra","user_age":37,"user_address":"Waterloo","user_postalcode":"7"}`)
	json.NewDecoder(rdr).Decode(&dUser)
	fmt.Println(dUser)
}
