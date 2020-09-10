package main

import (
	"backend-project/conf"
	"backend-project/http"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	h "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	runAPIServer := &h.Server{
		Handler: http.Router(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	runHealthCheck := &h.Server{
		Handler: http.RouterCheck(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	runValidateProduct := &h.Server{
		Handler: http.RouterValidate(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	var group errgroup.Group

	group.Go(func() error {
		return runAPIServer.ListenAndServe()
	})

	group.Go(func() error {
		return runHealthCheck.ListenAndServe()
	})

	group.Go(func() error {
		return runValidateProduct.ListenAndServe()
	})

	if err := group.Wait(); err != nil {
		fmt.Println("The server might be collided with another server")
	}

	// this method will be block graceful shutdown
	// if the listen address is working out
	go func() {
		if err := runHealthCheck.ListenAndServe(); err != nil && err == h.ErrServerClosed {
			log.Fatalf("The server listen to %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGSYS, syscall.SIGTERM)
	<-quit
	log.Println("There's an error occurred, shutting down the server")

	conf.ValidateDB()
}

