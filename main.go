package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/faridvaliyev1/v2/handlers"
)

func main() {
	//registering logger
	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	//creating handlers
	ph := handlers.NewProducts(l)

	//create new serve mux and register handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	//create new server
	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	//start the server
	go func() {
		l.Println("Starting the server on port 9090")

		err := s.ListenAndServe()

		if err != nil {
			l.Println("Error starting the server %s\n", err)
			os.Exit(1)
		}

	}()

	//shutting down the server

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	l.Println("Received terminate, gracefull shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(ctx)
}