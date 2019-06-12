package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"todo-api/database"
	"todo-api/utilities"
)

//InitStatusController Initializes status endpoints
func InitStatusController(r *mux.Router) {
	r.HandleFunc("/status", statusHandler).Methods("GET")
	r.HandleFunc("/status/db", dbHandler).Methods("GET")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	utilities.RespondOK(w)
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	err := database.DB.Ping()
	if err != nil {
		http.Error(w, "Error contacting DB: "+err.Error(), http.StatusNoContent)
	} else {
		w.Write([]byte("Database connection OK"))
	}
	utilities.RespondOK(w)
}
