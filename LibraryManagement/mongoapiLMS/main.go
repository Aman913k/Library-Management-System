package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aman913k/MONGOAPILMS/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Listening at port 5000...")
}
