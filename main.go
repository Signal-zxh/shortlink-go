package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	var idCounter uint64 = 0

	r.POST("/shorten", func(c *gin.Context) {
		var req struct {
			URL string `json:"url" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "请提供 url 字段"})
			return
		}

		idCounter++
		shortCode := fmt.Sprintf("%d", idCounter)

		// TODO: 生成短码、存数据库、返回短码
		c.JSON(http.StatusOK, gin.H{"short_code": shortCode})
	})

	r.Run(":8080")
}
