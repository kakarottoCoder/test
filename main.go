package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/user", GetUser)

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}
