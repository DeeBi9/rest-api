package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Deepanshuisjod/rest-api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/admin/information", handlers.EmployeeInformation)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	// Admin posting the project and rest details
	postRouter.HandleFunc("/admin/create", handlers.CreateProject)
	// Manager posting the tasks to the contributors
	postRouter.HandleFunc("/manager/{id:[0-9]+}", handlers.AssignTask)
	postRouter.Use(handlers.MiddlewareValidateID)

	s := &http.Server{
		Addr:         ":9091",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
