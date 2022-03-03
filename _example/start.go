package main

import (
	"fmt"
	"net/http"

	"github.com/gominima/mux"
)

func main() {
  rt := mux.NewRouter()

  rt.Get("/name/:id", func(w http.ResponseWriter, r *http.Request) {
	  param := rt.GetParam(r,"id")
	  fmt.Print(param)
	  w.Write([]byte("Hello"))
  })
  http.ListenAndServe(":3000", rt)
}