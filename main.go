package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/btcsuite/btcutil/base58"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var idCounter uint64 = 0

func main() {
	initDB()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/:shortcode", func(c *gin.Context) {
		shortCode := c.Param("shortcode")
		var longURL string
		err := db.QueryRow("SELECT long_url FROM shortlinks WHERE short_code = ?", shortCode).Scan(&longURL)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "短码不存在"})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
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

		_, err := db.Exec("INSERT INTO shortlinks (short_code, long_url) VALUES (?, ?)", shortCode, req.URL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"short_code": shortCode})
	})

	r.Run(":8081")
}

func initDB() {
	var err error
	dsn := "root:123456@tcp(localhost:3306)/shortlink?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS shortlinks (
		id INT AUTO_INCREMENT PRIMARY KEY,
		short_code VARCHAR(20) NOT NULL UNIQUE,
		long_url TEXT NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

	var maxID uint64
	db.QueryRow("SELECT COALESECE(MAX(id), 0) FROM shortlinks").Scan(&maxID)
	idCounter = maxID
}
