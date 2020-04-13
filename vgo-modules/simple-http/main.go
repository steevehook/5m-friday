package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Index(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	logger, _ := zap.NewProduction()
	logger.Info("successfully performed http request")
	logger.Info(quote.Hello())
	logger.Info(quoteV3.HelloV3())
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
