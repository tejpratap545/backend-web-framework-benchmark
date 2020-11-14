package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"language":  "go",
			"framework": "gin",
		})
	})

	// router runs on :8080 by default
	router.Run(":8001")
}
