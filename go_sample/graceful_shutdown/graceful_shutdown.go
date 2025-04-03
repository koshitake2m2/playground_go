package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// NOTE
// This is a simple HTTP server that demonstrates graceful shutdown.
// 1. Access `http://localhost:8080/`.
// 2. Press Ctrl+C to stop the server before the response is sent.
func helloHandler(w http.ResponseWriter, r *http.Request) {

	// 3. The server will wait for sending the response before shutting down.
	time.Sleep(4 * time.Second)

	// 3. The server will shut down before sending the response.
	// time.Sleep(10 * time.Second)

	w.Write([]byte("Hello, World!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Watch for interrupt signals (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Start the server in a goroutine
	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Wait for a signal
	<-stop
	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
