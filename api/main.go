package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	h := &handler{
		db: newDBClient(),
	}

	router := mux.NewRouter()
	router.HandleFunc("/ping", h.ping).Methods("GET")
	router.HandleFunc("/officer/{badge}", h.getOfficerByBadge).Methods("GET")

	port := os.Getenv("PORT")
	fmt.Println("starting server on port", port)
	http.ListenAndServe(":"+port, router)
}

type handler struct {
	db databaseInterface
}

func (h *handler) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("🏓 P O N G 🏓"))
}

func (h *handler) getOfficerByBadge(w http.ResponseWriter, r *http.Request) {
	ofc, err := h.db.getOfficerByBadge(mux.Vars(r)["badge"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ofc)
}
