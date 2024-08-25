package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	// TODO add option to helper to return HTML errors here
	baby, err := h.helperGetBaby(c, uid)
	if err != nil {
		return
	}
	words, err := h.db.ListWords(baby.Id)
	if err != nil {
		HandleErrorHTML(c, err)
		return
	}

	c.HTML(http.StatusOK, "list.tmpl", gin.H{
		"baby":  baby,
		"words": words,
	})
}

// POST /words/:uid/add payload: word=foo
func (h *Handler) HandleWordAdd(c *gin.Context) {
	uid := c.Param("uid")
	word := c.PostForm("word")

	learnedDate, isSet := c.GetPostForm("learned_date")
	if learnedDate == "" || !isSet {
		learnedDate = UTCTodayString()
	}

	baby, err := h.helperGetBaby(c, uid)
	if err != nil {
		return
	}

	wordId, err := h.db.AddWord(baby.Id, word, learnedDate)
	if err != nil {
		HandleErrorJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true, "word": gin.H{
			"id": wordId, "word": word,
		},
	})
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

func UTCTodayString() string {
	return time.Now().Format("2006-01-02")
}
