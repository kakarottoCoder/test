package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Get User Handler")
	io.WriteString(w, "hello,Git！")
	io.WriteString(w, "初始化")
	io.WriteString(w, "AAA")

}
