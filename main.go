package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Head("https://bit.ly/joinJungleDevs")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println(res.Request.URL)
	fmt.Println(res.Request.Header)
}
