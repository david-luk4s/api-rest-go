package routes

import (
	"log"
	"net/http"

	"github.com/david-luk4s/go-api-rest/controllers"
	"github.com/david-luk4s/go-api-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRouter() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleWare)
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.GetAllPersonality).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
