package main

import (
	"log"
	"os"

	"github.com/ab/baby-words/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := server.InitRouter()
	// Listen and serve on 0.0.0.0:8080
	log.Printf("baby-words/gin listening on port %v\n", port)
	r.Run(":" + port)
}
