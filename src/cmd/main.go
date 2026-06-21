package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/wowlikon/surl/src/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err == os.ErrNotExist {
		config.Save(config.Default())
		log.Println("Config created, please change secret and restart")
		return
	}
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	router.Run(":8080")
}
