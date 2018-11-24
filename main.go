package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ketabdoozak/backend/pkg/db"
	"log"
	"os"
)

func main() {
	err := db.DB().DB().Ping()
	if err != nil {
		log.Panic(err)
	}

	env := os.Getenv("ENV")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
		db.DB().LogMode(false)
	} else if env == "testing" {
		gin.SetMode(gin.TestMode)
		db.DB().LogMode(true)
	} else {
		gin.SetMode(gin.DebugMode)
		db.DB().LogMode(true)
	}

	r := gin.New()

	RegisterRoutes(r.Group("/"))

	err = r.Run(os.Getenv("SERVER_ADDRESS"))
	if err != nil {
		log.Panic(err)
	}

	err = db.Close()
	if err != nil {
		log.Panic(err)
	}
}
