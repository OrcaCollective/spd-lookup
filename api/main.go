package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	h := &handler{}

	router := mux.NewRouter()
	router.HandleFunc("/hello", h.hello).Methods("GET")

	port := os.Getenv("PORT")
	fmt.Println("starting server on port", port)
	http.ListenAndServe(":"+port, router)
}

type handler struct{}

func (h *handler) hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	message := "hello " + r.URL.Query().Get("name")
	w.Write([]byte(message))
}
