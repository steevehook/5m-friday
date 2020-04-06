package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func Index(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	logger, _ := zap.NewProduction()
	logger.Info("successfully performed http request")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
