package loan_service_go

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
	listener *listeners.Consumer
	server   *http.Server
}

func NewApplication(listener *listeners.Consumer, server *http.Server) *Application {
	return &Application{
		listener: listener,
		server:   server,
	}
}

func (app *Application) Start() error {
	// Inicia os dois serviços em goroutines separadas
	errChan := make(chan error, 2)

	// Inicia o listener Kafka
	go func() {
		if err := app.listener.Start(); err != nil {
			errChan <- fmt.Errorf("kafka listener error: %v", err)
		}
	}()

	// Inicia o servidor HTTP
	go func() {
		if err := app.server.ListenAndServe(); err != http.ErrServerClosed {
			errChan <- fmt.Errorf("http server error: %v", err)
		}
	}()

	return <-errChan
}

func (app *Application) Close() error {
	var errors []error

	// Fecha o servidor HTTP gracefully
	if err := app.server.Shutdown(context.Background()); err != nil {
		errors = append(errors, fmt.Errorf("error shutting down http server: %v", err))
	}

	// Fecha o listener Kafka
	if err := app.listener.Close(); err != nil {
		errors = append(errors, fmt.Errorf("error closing kafka listener: %v", err))
	}

	if len(errors) > 0 {
		return fmt.Errorf("shutdown errors: %v", errors)
	}
	return nil
}

func provideHTTPServer(service *services.ValidationService) *http.Server {
	router := chi.NewRouter()

	router.Post("/validate-card", func(w http.ResponseWriter, r *http.Request) {
		// sua lógica de handler aqui
	})

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}

// main.go
func main() {
	app, err := InitializeApp()
	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Start(); err != nil {
			log.Printf("Error starting application: %v", err)
			sigChan <- syscall.SIGTERM
		}
	}()

	sig := <-sigChan
	log.Printf("Received signal: %v", sig)

	if err := app.Close(); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
}
