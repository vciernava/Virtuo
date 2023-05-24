package cmd

import (
	"fmt"
	"github.com/vciernava/virtuo/router"
	"log"
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
	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}
