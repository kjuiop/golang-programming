package main

import (
	"golang-programming/http"
	"golang-programming/util"
	"log"
)

type Handler struct {
	cfg        *util.Config
	httpClient *http.HttpClient
}

func NewHandler() (*Handler, error) {
	h := new(Handler)

	var err error

	h.cfg, err = util.ConfInitialize()
	if err != nil {
		log.Println("[NewHandler] failed config initialize : ", err)
		return nil, err
	}

	h.httpClient, err = http.HttpClientInitialize()
	if err != nil {
		log.Println("[NewHandler] failed zookeeper initialize : ", err)
	}

	return h, err
}

func (h *Handler) Close() {
	log.Println("[main] Handler Close")
}
