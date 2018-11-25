package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ketabdoozak/backend/pkg/db"
	"log"
	"net/http"
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

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:    []string{"Origin", "Content-Type", "User-Agent", "User-Token"},
	}))

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
