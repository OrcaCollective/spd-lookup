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
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return
	}
}

// NewRouter is the router constructor
func NewRouter(h handler.Interface) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/ping", h.Ping).Methods("GET")
	router.HandleFunc("/departments", h.DescribeDepartments).Methods("GET")

	router.HandleFunc("/seattle/metadata", h.SeattleOfficerMetadata).Methods("GET")
	router.HandleFunc("/seattle/officer", h.SeattleStrictMatch).Methods("GET")
	router.HandleFunc("/seattle/officer/search", h.SeattleFuzzySearch).Methods("GET")
	router.HandleFunc("/seattle/officer/historical", h.SeattleStrictMatchHistorical).Methods("GET")

	router.HandleFunc("/tacoma/metadata", h.TacomaOfficerMetadata).Methods("GET")
	router.HandleFunc("/tacoma/officer", h.TacomaStrictMatch).Methods("GET")
	router.HandleFunc("/tacoma/officer/search", h.TacomaFuzzySearch).Methods("GET")

	router.HandleFunc("/portland/metadata", h.PortlandOfficerMetadata).Methods("GET")
	router.HandleFunc("/portland/officer", h.PortlandStrictMatch).Methods("GET")
	router.HandleFunc("/portland/officer/search", h.PortlandFuzzySearch).Methods("GET")

	router.HandleFunc("/auburn/metadata", h.AuburnOfficerMetadata).Methods("GET")
	router.HandleFunc("/auburn/officer", h.AuburnStrictMatch).Methods("GET")
	router.HandleFunc("/auburn/officer/search", h.AuburnFuzzySearch).Methods("GET")
	return router
}
