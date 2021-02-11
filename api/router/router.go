package router

import (
	"fmt"
	"net/http"
	"os"

	"spd-lookup/api/data"
	"spd-lookup/api/handler"

	"github.com/gorilla/mux"
)

// Start starts up the router
func Start() {
	router := NewRouter(handler.NewHandler(data.NewClient(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)))

	port := os.Getenv("PORT")
	fmt.Println("starting server on port", port)
	http.ListenAndServe(":"+port, router)
}

// NewRouter is the router constructor
func NewRouter(h handler.HandInterface) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/ping", h.Ping).Methods("GET")
	router.HandleFunc("/departments", h.DescribeDepartments).Methods("GET")

	router.HandleFunc("/seattle/metadata", h.SeattleOfficerMetadata).Methods("GET")
	router.HandleFunc("/seattle/officer", h.SeattleStrictMatch).Methods("GET")
	router.HandleFunc("/seattle/officer/search", h.SeattleFuzzySearch).Methods("GET")

	router.HandleFunc("/tacoma/metadata", h.TacomaOfficerMetadata).Methods("GET")
	router.HandleFunc("/tacoma/officer", h.TacomaStrictMatch).Methods("GET")
	router.HandleFunc("/tacoma/officer/search", h.TacomaFuzzySearch).Methods("GET")
	return router
}
