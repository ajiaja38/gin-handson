package main

import (
	"os"
	"regexp"
	"res-gin/src/config"
	"res-gin/src/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	ginMode := os.Getenv("GIN_MODE")

	if ginMode != "" {
		gin.SetMode(ginMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	port := ":" + os.Getenv("PORT")
	globalPrefix := "api/v1"

	db, err := config.DBConnection()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.SetTrustedProxies([]string{"103.150.190.175", "127.0.0.1"})
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			matched, _ := regexp.MatchString(`^(https?://)?([a-zA-Z0-9.-]+\.lskk\.co\.id|[a-zA-Z0-9.-]+\.pptik\.id|localhost(:[0-9]+)?)$`, origin)
			return matched
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group(globalPrefix)
	{
		api.GET("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"code":    200,
				"status":  true,
				"message": "Welcome to Gin API",
			})
		})
	}

	router.SetupRouter(api, db)

	log.Infof("🚀 Application listening on http://localhost%s/%s", port, globalPrefix)

	r.Run(port)
}
