package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	calendar "github.com/cheyzie/go_calendar"
	"github.com/cheyzie/go_calendar/pkg/handler"
	"github.com/cheyzie/go_calendar/pkg/repository"
	"github.com/cheyzie/go_calendar/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Calendar API
// @version 0.9

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("error in config initialization: %s", err)
	}
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("error in loading env variables: %s", err)
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed db connection: %s", err)
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	server := new(calendar.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error caused on server startup: %s", err)
		}
	}()
	logrus.Debug("Server is started...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Debug("Server is shutting down...")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error caused on server shutting down: %s", err)
	}
	if err := db.Close(); err != nil {
		logrus.Fatalf("error caused on db conection closing: %s", err)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
