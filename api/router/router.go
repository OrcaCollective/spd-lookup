package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/OrcaCollective/spd-lookup/api/data"
	"github.com/OrcaCollective/spd-lookup/api/handler"

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
	if port == "" {
		port = "1312"
	}
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

	router.HandleFunc("/seattle-wa/metadata", h.SeattleOfficerMetadata).Methods("GET")
	router.HandleFunc("/seattle-wa/officer", h.SeattleStrictMatch).Methods("GET")
	router.HandleFunc("/seattle-wa/officer/search", h.SeattleFuzzySearch).Methods("GET")
	router.HandleFunc("/seattle-wa/officer/historical", h.SeattleStrictMatchHistorical).Methods("GET")

	router.HandleFunc("/tacoma-wa/metadata", h.TacomaOfficerMetadata).Methods("GET")
	router.HandleFunc("/tacoma-wa/officer", h.TacomaStrictMatch).Methods("GET")
	router.HandleFunc("/tacoma-wa/officer/search", h.TacomaFuzzySearch).Methods("GET")

	router.HandleFunc("/portland-or/metadata", h.PortlandOfficerMetadata).Methods("GET")
	router.HandleFunc("/portland-or/officer", h.PortlandStrictMatch).Methods("GET")
	router.HandleFunc("/portland-or/officer/search", h.PortlandFuzzySearch).Methods("GET")

	router.HandleFunc("/auburn-wa/metadata", h.AuburnOfficerMetadata).Methods("GET")
	router.HandleFunc("/auburn-wa/officer", h.AuburnStrictMatch).Methods("GET")
	router.HandleFunc("/auburn-wa/officer/search", h.AuburnFuzzySearch).Methods("GET")

	router.HandleFunc("/lakewood-wa/metadata", h.LakewoodOfficerMetadata).Methods("GET")
	router.HandleFunc("/lakewood-wa/officer", h.LakewoodStrictMatch).Methods("GET")
	router.HandleFunc("/lakewood-wa/officer/search", h.LakewoodFuzzySearch).Methods("GET")

	router.HandleFunc("/bellevue-wa/metadata", h.BellevueOfficerMetadata).Methods("GET")
	router.HandleFunc("/bellevue-wa/officer", h.BellevueStrictMatch).Methods("GET")
	router.HandleFunc("/bellevue-wa/officer/search", h.BellevueFuzzySearch).Methods("GET")

	router.HandleFunc("/port_of_seattle-wa/metadata", h.PortOfSeattleOfficerMetadata).Methods("GET")
	router.HandleFunc("/port_of_seattle-wa/officer", h.PortOfSeattleStrictMatch).Methods("GET")
	router.HandleFunc("/port_of_seattle-wa/officer/search", h.PortOfSeattleFuzzySearch).Methods("GET")

	router.HandleFunc("/thurston_county-wa/metadata", h.ThurstonCountyOfficerMetadata).Methods("GET")
	router.HandleFunc("/thurston_county-wa/officer", h.ThurstonCountyStrictMatch).Methods("GET")
	router.HandleFunc("/thurston_county-wa/officer/search", h.ThurstonCountyFuzzySearch).Methods("GET")

	router.HandleFunc("/renton-wa/metadata", h.RentonOfficerMetadata).Methods("GET")
	router.HandleFunc("/renton-wa/officer", h.RentonStrictMatch).Methods("GET")
	router.HandleFunc("/renton-wa/officer/search", h.RentonFuzzySearch).Methods("GET")

	router.HandleFunc("/olympia-wa/metadata", h.OlympiaOfficerMetadata).Methods("GET")
	router.HandleFunc("/olympia-wa/officer", h.OlympiaStrictMatch).Methods("GET")
	router.HandleFunc("/olympia-wa/officer/search", h.OlympiaFuzzySearch).Methods("GET")

	// Backwards compatability routes, don't change
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

	router.HandleFunc("/lakewood/metadata", h.LakewoodOfficerMetadata).Methods("GET")
	router.HandleFunc("/lakewood/officer", h.LakewoodStrictMatch).Methods("GET")
	router.HandleFunc("/lakewood/officer/search", h.LakewoodFuzzySearch).Methods("GET")

	router.HandleFunc("/bellevue/metadata", h.BellevueOfficerMetadata).Methods("GET")
	router.HandleFunc("/bellevue/officer", h.BellevueStrictMatch).Methods("GET")
	router.HandleFunc("/bellevue/officer/search", h.BellevueFuzzySearch).Methods("GET")

	router.HandleFunc("/port_of_seattle/metadata", h.PortOfSeattleOfficerMetadata).Methods("GET")
	router.HandleFunc("/port_of_seattle/officer", h.PortOfSeattleStrictMatch).Methods("GET")
	router.HandleFunc("/port_of_seattle/officer/search", h.PortOfSeattleFuzzySearch).Methods("GET")

	router.HandleFunc("/thurston_county/metadata", h.ThurstonCountyOfficerMetadata).Methods("GET")
	router.HandleFunc("/thurston_county/officer", h.ThurstonCountyStrictMatch).Methods("GET")
	router.HandleFunc("/thurston_county/officer/search", h.ThurstonCountyFuzzySearch).Methods("GET")

	router.HandleFunc("/renton/metadata", h.RentonOfficerMetadata).Methods("GET")
	router.HandleFunc("/renton/officer", h.RentonStrictMatch).Methods("GET")
	router.HandleFunc("/renton/officer/search", h.RentonFuzzySearch).Methods("GET")

	router.HandleFunc("/olympia/metadata", h.OlympiaOfficerMetadata).Methods("GET")
	router.HandleFunc("/olympia/officer", h.OlympiaStrictMatch).Methods("GET")
	router.HandleFunc("/olympia/officer/search", h.OlympiaFuzzySearch).Methods("GET")
	return router
}
