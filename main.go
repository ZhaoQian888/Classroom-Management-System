package main

import (
	"GoPratice/database"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	err := database.Init()
	fmt.Print(err)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.String(200, "fdfsdfdsfdsfsdf")

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
