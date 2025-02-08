package main

import (
	"github.com/fabriciolfj/loan-service-go/controller"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func appHttp() {
	mux := http.NewServeMux()

	pc, _ := InitializeProductControllerWire()
	http.HandleFunc("/products", controller.RecoveryMiddleware(pc.HandleProduct))

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	go func() {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Log.Fatal("fail star server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Log.Fatal("Server forced to shutdown: ", err)
	}

	log.Log.Info("Server exiting")
}

func listenerProcessLoan() {
	app, err := InitListenerProcessLoan()

	if err != nil {
		panic(err)
	}

	igChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	errChan := make(chan error, 2)

	go func() {
		if err := listener.Start(); err != nil {
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
	listenerProcessLoan()
}
