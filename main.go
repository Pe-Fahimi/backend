package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ketabdoozak/backend/pkg/db"
	"log"
)

func main() {
	err := db.DB().DB().Ping()
	if err != nil {
		log.Panic(err)
	}

	r := gin.New()

	RegisterRoutes(r.Group("/"))

	err = r.Run("0.0.0.0:8080") // TODO: Get from env
	if err != nil {
		log.Panic(err)
	}

	err = db.Close()
	if err != nil {
		log.Panic(err)
	}
}
