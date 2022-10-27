package main

import (
	"golang-programming/http"
	"golang-programming/logger"
	"golang-programming/util"
	"log"
)

type Handler struct {
	cfg        *util.Config
	log        *logger.Logrus
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

	h.log, err = logger.LogInitialize(h.cfg.LogInfo.LogPath, h.cfg.LogInfo.LogLevel)
	if err != nil {
		log.Println("[NewHandler] failed log initialize : ", err)
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
