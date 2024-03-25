package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/ab/baby-words/db"
)

// GET /
func HandleRoot(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"success": true})
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"region":   os.Getenv("FLY_REGION"),
		"clientIP": c.ClientIP(),
	})
}

// GET /words/:uid/list
func HandleWordList(c *gin.Context) {
	uid := c.Param("uid")

	baby, err := helperGetBaby(c, uid)
	if err != nil {
		return
	}
	words := baby.ListWords()
	c.HTML(http.StatusOK, "list.tmpl", gin.H{
		"baby":  baby,
		"words": words,
	})
}

// GET /words/:uid/check/:word
func HandleWordCheck(c *gin.Context) {
	uid := c.Param("uid")
	word := c.Param("word")

	baby, err := helperGetBaby(c, uid)
	if err != nil {
		return
	}
	wordInfo, err := baby.WordInfo(word)
	if err != nil {
		HandleErrorJSON(c, err)
	}

	c.JSON(http.StatusOK, gin.H{"TODO": "return word info", "wordInfo": wordInfo})
}

// POST /words/:uid/add/:word
func HandleWordAdd(c *gin.Context) {
	uid := c.Param("uid")
	word := c.Param("word")

	baby, err := helperGetBaby(c, uid)
	if err != nil {
		return
	}

	wordInfo, err := baby.AddWord(word)
	if err != nil {
		HandleErrorJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "TODO": "return word info", "wordInfo": wordInfo})
}

func helperGetBaby(c *gin.Context, uid string) (*db.BabyStruct, error) {
	baby, err := db.GetBaby(uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No such baby"})
		return nil, fmt.Errorf("No such baby")
	}

	return baby, nil
}

func HandleErrorJSON(c *gin.Context, err error) {
	c.JSON(
		http.StatusInternalServerError,
		gin.H{"error": true, "success": false, "message": fmt.Sprintf("%v", err)},
	)
}

func HandleErrorHTML(c *gin.Context, err error) {
	c.HTML(
		http.StatusInternalServerError,
		"error.tmpl",
		gin.H{"error": true, "message": fmt.Sprintf("%v", err)},
	)
}
