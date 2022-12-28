package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type welcomeResponse struct {
	Message string `json:"message"`
}

func hello(c *gin.Context) {
	var response = welcomeResponse{Message: "Hello World! There"}
	c.IndentedJSON(http.StatusOK, response)
}

func setupServer() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", hello)
	return router
}

func main() {
	if err := setupServer().Run(); err != nil {
		log.Fatalf("Error while starting server!")
	}

}
