package main

import (
	"net/http"
	"strconv"

	"github.com/btcsuite/btcutil/base58"
	"github.com/gin-gonic/gin"
)

var urlMap = make(map[string]string)
var idCounter uint64 = 0

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/:shortcode", func(c *gin.Context) {
		shortCode := c.Param("shortcode")
		longURL, exists := urlMap[shortCode]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "短码不存在"})
			return
		}
		c.Redirect(http.StatusMovedPermanently, longURL)
	})

	r.POST("/shorten", func(c *gin.Context) {
		var req struct {
			URL string `json:"url" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "请提供 url 字段"})
			return
		}

		idCounter++
		shortCode := base58.Encode([]byte(strconv.FormatUint(idCounter, 10)))
		urlMap[shortCode] = req.URL

		// TODO: 生成短码、存数据库、返回短码
		c.JSON(http.StatusOK, gin.H{"short_code": shortCode})
	})

	r.Run(":8080")
}
