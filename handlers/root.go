package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleRoot(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"success": true})
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"region":   os.Getenv("FLY_REGION"),
		"clientIP": c.ClientIP(),
	})
}
