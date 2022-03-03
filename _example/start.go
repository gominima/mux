package main

import (
	"net/http"

	"github.com/gominima/mux"
)

func main() {
	mux.Abc(func(r http.Response, req *http.Request) {
		
	})
}