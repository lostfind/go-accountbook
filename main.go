package main

import (
	"go-accountbook/registry"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	// historyCtr := registry.InjectionHistoryCtr()
	godotenv.Load()

	historyCtl := registry.InjectionHistory()

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {

		log.WithFields(log.Fields{
			"animal": "walrus",
			"size":   10,
		}).Info("A group of walrus emerges from the ocean")

		contextLogger := log.WithFields(log.Fields{
			"common": "this is a common field",
			"other":  "I also should be logged always",
		})

		contextLogger.Info("I'll be logged with common and other field")
		contextLogger.Info("Me too")

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Use(cors.Default())
	server.GET("/history/:id", historyCtl.GetHistory)
	server.GET("/history", historyCtl.GetHistories)
	server.POST("/assetmove", historyCtl.MoveAsset)

	server.Run()
}
