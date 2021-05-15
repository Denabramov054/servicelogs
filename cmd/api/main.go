package main

import (
	"context"
	"os"
	"os/signal"
	"project/servicelogs/pkg/controller"
	"project/servicelogs/pkg/repository"
	"project/servicelogs/pkg/service"
	"syscall"

	servicelogs "project/servicelogs"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf(err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DBNAME"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	})

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	controller := controller.NewController(services)

	srv := new(servicelogs.Server)

	go func() {
		if err := srv.Run(os.Getenv("PORT"), controller.InitRoutes()); err != nil {
			logrus.Fatalf(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
