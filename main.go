package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	cors "github.com/efuchsman/EliFuchsmanBE/config"
	"github.com/efuchsman/EliFuchsmanBE/handlers"
	elifuchsman "github.com/efuchsman/EliFuchsmanBE/internal/eli_fuchsman"
	elifuchsmandb "github.com/efuchsman/EliFuchsmanBE/internal/eli_fuchsman_db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	log.Info("Starting Eli's Backend Application")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	awsRegion := os.Getenv("AWS_REGION")
	log.Infof("AWS Region: %s", awsRegion)

	edb, err := elifuchsmandb.NewEliFuchsmanDB(awsRegion, "")
	if err != nil {
		log.WithError(err).Fatal("Failure opening database connection")
	}

	eliClient := elifuchsman.NewEliFuchsmanClient(edb)
	tableOne := os.Getenv("DYNAMO_TABLE_1")
	log.Infof("Using DynamoDB Table: %s", tableOne)

	router := mux.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "tableOne", tableOne)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Eli's Backend")
	})

	eliHandler := handlers.NewHandler(eliClient)
	router.HandleFunc("/info", eliHandler.GetBasicInfo).Methods("GET")

	handler := cors.SetCORSHeader(router)

	port := 8000
	fmt.Printf("Server is running on :%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	fmt.Println("Application started successfully")
}
