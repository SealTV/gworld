package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"os/signal"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	pflag.String("pong_message", "pong", "Pong message")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	log.Println("service started")
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("service stopped")
}

func run(ctx context.Context) error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", pingHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("failed to shutdown server: %v", err)
		}
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("error on runing server: %w", err)
	}

	return nil
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	a, b := rand.Int(), rand.Int()

	msg := fmt.Sprintf("%s %d", viper.GetString("pong_message"), sum(a, b))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func sum(a, b int) int {
	return a + b
}
