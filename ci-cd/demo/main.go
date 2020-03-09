package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/greetings", func(c *gin.Context) {
		person := c.Query("person")
		helloGreetings := greetings(person)
		c.JSON(200, gin.H{
			"message": helloGreetings,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func greetings(person string) string {
	if person == "" {
		person = "World"
	}

	return fmt.Sprintf("Hello, %s", person)
}
