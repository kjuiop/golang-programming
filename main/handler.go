package main

import (
	"golang-programming/http"
	"log"
)

type Handler struct {
	httpClient *http.HttpClient
}

func NewHandler() (*Handler, error) {
	h := new(Handler)

	var err error
	h.httpClient, err = http.HttpClientInitialize()
	if err != nil {
		log.Println("[NewHandler] failed zookeeper initialize : ", err)
	}

	return h, err
}

func (h *Handler) Close() {
	log.Println("[main] Handler Close")
}
