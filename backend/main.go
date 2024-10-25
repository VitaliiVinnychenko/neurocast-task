package main

import (
	"context"
	"github.com/VitaliiVinnychenko/neurocast-task/backend/routes"
	"github.com/VitaliiVinnychenko/neurocast-task/backend/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title GoLang Rest API Notes
// @version 1.0
// @description GoLang - Gin - RESTful - MongoDB - Redis
// @termsOfService https://swagger.io/terms/

// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	services.LoadConfig()
	services.InitMongoDB()

	routes.InitGin()
	router := routes.New()

	server := &http.Server{
		Addr:         services.Config.ServerAddr + ":" + services.Config.ServerPort,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 30,
		Handler:      router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 15 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
