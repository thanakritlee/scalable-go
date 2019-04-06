package main

import (
	"log"
	"net/http"
	"os"

	"github.com/thanakritlee/scalable-go/entryserver/router"
)

func main() {

	router := router.GetRouter()

	port := os.Getenv("ENTRYSERVER_PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("http server started on :%s\n", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
