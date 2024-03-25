package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"

	"github.com/ab/baby-words/handlers"
)

func InitRouter() *gin.Engine {
	// gin.DisableConsoleColor()  // Disable Console Color
	r := gin.Default()

	r.LoadHTMLGlob(dirOfExecutable() + "/templates/*.tmpl")

	// Never use X-Forwarded-For
	r.SetTrustedProxies(nil)

	// Get client IP from fly.io header
	// https://fly.io/docs/reference/runtime-environment/#fly-client-ip
	r.TrustedPlatform = "Fly-Client-IP"

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"healthy": true, "region": os.Getenv("FLY_REGION")})
	})

	// Root default page
	r.GET("/", handlers.HandleRoot)

	// List words
	r.GET("/words/:uid/list", handlers.HandleWordList)

	// word get / create
	r.GET("/words/:uid/check/:word", handlers.HandleWordCheck)
	r.POST("/words/:uid/add/:word", handlers.HandleWordAdd)

	return r
}

// Get the directory of the currently running executable
func dirOfExecutable() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return filepath.Dir(ex)
}

func connectDatabase(filename string) sqlx.DB {
	db, err := sqlx.Connect("sqlite", filename)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
