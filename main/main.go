package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello golang-programming")

	h, err := NewHandler()
	if err != nil {
		log.Println("[main] failed NewHandler : ", err)
		os.Exit(1)
	}
	defer h.Close()

	str, _, _ := h.httpClient.Get("http://localhost:3010/api/transcoder/v2/hello/test", 3)
	fmt.Println(str)
}
