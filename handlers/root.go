package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/ab/baby-words/db"
)

type Handler struct {
	db *db.Connection
}

func NewHandler(sqliteFilename string) *Handler {
	conn := db.NewConnection(sqliteFilename)

	return &Handler{
		db: conn,
	}
}

// GET /
func (h *Handler) HandleRoot(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"success": true})
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"region":   os.Getenv("FLY_REGION"),
		"clientIP": c.ClientIP(),
	})
}

// POST
/*
**babies**

- id
- slug
- name
- birth_date
- created_at
*/
func (h *Handler) HandleCreateBaby(c *gin.Context) {
	name := c.PostForm("name")
	birthdayString := c.PostForm("birth_date")

	baby, err := h.db.CreateBaby(name, birthdayString)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create baby"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"baby": gin.H{
			"id":   baby.Id,
			"slug": baby.Slug,
		},
	})
}

// GET /words/:uid/list
func (h *Handler) HandleWordList(c *gin.Context) {
	uid := c.Param("uid")

	baby, err := h.helperGetBaby(c, uid)
	if err != nil {
		return
	}
	words := baby.ListWords()
	c.HTML(http.StatusOK, "list.tmpl", gin.H{
		"baby":  baby,
		"words": words,
	})
}

// GET /words/:uid/check?word=foo
func (h *Handler) HandleWordCheck(c *gin.Context) {
	uid := c.Param("uid")
	word := c.Query("word")

	baby, err := h.helperGetBaby(c, uid)
	if err != nil {
		return
	}
	wordInfo, err := baby.WordInfo(word)
	if err != nil {
		HandleErrorJSON(c, err)
	}

	c.JSON(http.StatusOK, gin.H{"TODO": "return word info", "wordInfo": wordInfo})
}

// POST /words/:uid/add payload: word=foo
func (h *Handler) HandleWordAdd(c *gin.Context) {
	uid := c.Param("uid")
	word := c.PostForm("word")

	baby, err := h.helperGetBaby(c, uid)
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

func (h *Handler) helperGetBaby(c *gin.Context, slug string) (*db.BabyStruct, error) {
	baby, err := h.db.GetBaby(slug)
	if err != nil {
		log.Printf("Failed to find baby slug=%v: %v", slug, err)
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
