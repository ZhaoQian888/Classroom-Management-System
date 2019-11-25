package main

import "github.com/gin-gonic/gin"
import "database"

func main() {
	database.Init()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.String(200, "fdfsdfdsfdsfsdf")

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
