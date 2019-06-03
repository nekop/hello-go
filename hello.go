package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Gin\n",)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

