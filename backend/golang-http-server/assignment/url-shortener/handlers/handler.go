package handlers

import (
	"fmt"
	"net/http"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) URLHandler {
	return URLHandler{
		repo: repo,
	}
}

func (h *URLHandler) Get(c *gin.Context) {
	// TODO: answer here
	param := c.Param("url")
	data, err := h.repo.Get(param)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "url not found",
		})
		return
	}
	fmt.Println("MYDATA:", data.LongURL)
	c.Header("Location", data.LongURL)
	c.JSON(http.StatusFound, gin.H{
		"Data": data,
	})

}

func (h *URLHandler) Create(c *gin.Context) {
	// TODO: answer here
	if c.Request.Method == http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	var url entity.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	data, err := h.repo.Create(url.LongURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Data": data,
	})
}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	// TODO: answer here
	if c.Request.Method == http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	var url entity.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	data, err := h.repo.CreateCustom(url.LongURL, url.ShortURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": data,
	})
}
