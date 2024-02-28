package router

import (
	"github.com/Aman913k/MONGOAPILMS/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/books", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/api/book", controller.CreateBook).Methods("POST")
	router.HandleFunc("/api/book/{id}", controller.MarkAsAvailable).Methods("PUT")
	router.HandleFunc("/api/book/{id}", controller.DeleteBook).Methods("DELETE")
	router.HandleFunc("/api/deleteallbook", controller.DeleteAllBooks).Methods("DELETE")

	return router
}
