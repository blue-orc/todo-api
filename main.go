package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"todo-api/controllers"
	"todo-api/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	database.Initialize()
	defer database.DB.Close()

	r := mux.NewRouter()
	initializeControllers(r)
	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))
}

func initializeControllers(r *mux.Router) {
	controllers.InitStatusController(r)
	controllers.InitTodoController(r)
}
