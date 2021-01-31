package router

import (
	"fmt"
	"net/http"
	"os"

	"spd-lookup/api/handler"

	"github.com/gorilla/mux"
)

// Start starts up the router
func Start() {
	h := handler.NewHandler()

	router := mux.NewRouter()
	router.HandleFunc("/ping", h.Ping).Methods("GET")
	router.HandleFunc("/departments", h.DescribeDepartments).Methods("GET")

	router.HandleFunc("/seattle/metadata", h.SeattleOfficerMetadata).Methods("GET")
	router.HandleFunc("/seattle/officer", h.SeattleStrictMatch).Methods("GET")
	router.HandleFunc("/seattle/officer/search", h.SeattleFuzzySearch).Methods("GET")

	router.HandleFunc("/tacoma/metadata", h.TacomaOfficerMetadata).Methods("GET")
	router.HandleFunc("/tacoma/officer", h.TacomaStrictMatch).Methods("GET")
	router.HandleFunc("/tacoma/officer/search", h.TacomaFuzzySearch).Methods("GET")

	port := os.Getenv("PORT")
	fmt.Println("starting server on port", port)
	http.ListenAndServe(":"+port, router)
}
