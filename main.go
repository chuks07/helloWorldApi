package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main ( ){
	router:= gin.Default()

	router.GET("/chuks", helloWorldhandler)
	port :=os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}


}

func helloWorldhandler(c *gin.Context) {
	c.JSON(200,gin.H{
		"message":"hello world",
		"status": "we are live",
		"age": "25",
		"name": "father ruben",
	})
}