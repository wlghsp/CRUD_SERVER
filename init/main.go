package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloworld)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("에러가 떠버렸는데요??")
		panic(err)
	}
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}
