package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
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

	res, err := h.httpClient.GetRequest("http://localhost:3010/api/transcoder/v2/hello/test", 3)
	if err != nil {
		fmt.Println(err)
	}
	prefix := h.log.WithFields(logrus.Fields{})
	prefix.Data["url"] = "http://localhost:3010/api/transcoder/v2/hello/test"
	prefix.Data["response"] = res
	prefix.Data["connection_time_out"] = 3

	h.log.Info(prefix, "http get test")
}
