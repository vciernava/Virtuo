package cmd

import (
	"fmt"
	"github.com/vciernava/Virtuo/router"
	"log"
	"net/http"
	"os"
)

func Execute() {
	r := router.Configure()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
