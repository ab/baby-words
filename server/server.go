package server

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

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

	// Set up handler and connect to database
	dbFile := dirOfExecutable() + "/data/db.sqlite3"
	h := handlers.NewHandler(dbFile)

	// Root default page
	r.GET("/", h.HandleRoot)

	// Create baby
	r.POST("/baby", h.HandleCreateBaby)

	// List words
	r.GET("/words/:uid/list", h.HandleWordList)

	// word create
	r.POST("/words/:uid/add", h.HandleWordAdd)

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
