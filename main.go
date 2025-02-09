package main

import (
	"context"
	"github.com/fabriciolfj/loan-service-go/controller"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func appHttp() {
	pc, err := InitControllerLoan()
	if err != nil {
		log.Fatal("Failed to initialize controller:", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/loan", controller.RecoveryMiddleware(pc.HandlerLoan))

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("fail star server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Printf("Server exiting")
}

func listenerProcessLoan() {
	app, err := InitListenerProcessLoan()

	if err != nil {
		panic(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	errChan := make(chan error, 2)

	go func() {
		if err := app.Start(); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		log.Printf("Erro: %v", err)
	case sig := <-sigChan:
		log.Printf("Sinal recebido: %v", sig)
	}

	if err := app.Close(); err != nil {
		log.Printf("Erro ao fechar listener1: %v", err)
	}

}

func main() {
	go listenerProcessLoan()
	appHttp()
}
