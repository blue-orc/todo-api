package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"todo-api/database/repositories"
	"todo-api/models"
	"todo-api/utilities"
)

//InitTodoController Initializes Todo endpoints
func InitTodoController(r *mux.Router) {
	r.HandleFunc("/todo", selectHandler).Methods("GET")
	r.HandleFunc("/todo", insertHandler).Methods("POST")
	r.HandleFunc("/todo", deleteHandler).Methods("DELETE")
}

func selectHandler(w http.ResponseWriter, r *http.Request) {
	un := r.URL.Query().Get("username")
	if un == "" {
		utilities.RespondBadRequest(w, "Username must be provided")
		return
	}
	tr := repositories.Todo{}
	ts, err := tr.SelectByUsername(un)
	if err != nil {
		utilities.RespondBadRequest(w, "Error communicating with database: "+err.Error())
	}
	utilities.RespondJSON(w, ts)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	tID, err := strconv.Atoi(r.URL.Query().Get("todoId"))
	if err != nil {
		utilities.RespondBadRequest(w, "TodoID must be integer value "+err.Error())
		return
	}
	tr := repositories.Todo{}
	tr.Delete(tID)
	utilities.RespondOK(w)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	t := models.Todo{}
	utilities.ReadJsonHttpBody(w, r, &t)
	tr := repositories.Todo{}
	tr.Insert(t)
	utilities.RespondOK(w)
}
