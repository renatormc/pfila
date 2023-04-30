package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reantormc/pfila/api/config"
	"github.com/reantormc/pfila/api/database"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.LoadConfig()
	database.Migrate()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.Run()
}
