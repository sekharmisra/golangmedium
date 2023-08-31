package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://www.google.com")
	data, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(data))
}
