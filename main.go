package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var router *gin.Engine

func InitRouter() {
	//  gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router = gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func main() {
	var numbers []int

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d  pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
	fmt.Println(time.Now())

	InitRouter()
	router.Run() // listen and serve on 0.0.0.0:8080

}
