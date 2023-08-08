package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/xbizzybone/go-sii-info/sii"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.Logger

func init() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	logConfig := zap.NewProductionConfig()
	logConfig.EncoderConfig.TimeKey = "timestamp"
	logConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	zapLogger, err = logConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	zapLogger = zapLogger.With(zap.String("service", "go-clean-code"))
	zapLogger.Info("Logger initialized")
}

func main() {
	defer zapLogger.Sync()

	zapLogger.Info("Starting server")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	sii.ApplyRoutes(app, zapLogger)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
