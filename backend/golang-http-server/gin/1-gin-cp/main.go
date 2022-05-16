package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dari contoh yang telah diberikan, buatlah HTTP server sederhana dengan menggunakan Gin.
// Buatlah server yang memiliki address localhost dengan port 3000.
// Buatlah route "/hello" yang menampilkan JSON {"message": "hello"} dan "/world" yang menampilkan JSON {"message": "world"}

func GetGinRoute() *gin.Engine {
	r := gin.Default()
	//{"message": "hello"}
	r.GET("/hello", func(c *gin.Context) {
		// req.resp
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	//{"message": "World"}
	r.GET("/world", func(c *gin.Context) {
		// req.resp
		c.JSON(http.StatusOK, gin.H{
			"message": "world",
		})
	})

	return r
	// return &gin.Engine{} // TODO: replace this
}

func main() {
	r := GetGinRoute()
	r.Run("localhost:3000")
}
