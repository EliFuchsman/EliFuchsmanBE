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

	log.Info("Starting Eli's Backend Application!")
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
	tableTwo := os.Getenv("DYNAMO_TABLE_2")
	tableThree := os.Getenv("DYNAMO_TABLE_3")
	bucket := os.Getenv("S3_BUCKET")
	bucketKey := os.Getenv("S3_BUCKET_KEY")

	router := mux.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "TableOne", tableOne)
			infoFile := "data/basic_info.json"
			ctx = context.WithValue(ctx, "InfoFile", infoFile)
			summaryFile := "data/summary.json"
			ctx = context.WithValue(ctx, "SummaryFile", summaryFile)
			ctx = context.WithValue(ctx, "TableTwo", tableTwo)
			ctx = context.WithValue(ctx, "Bucket", bucket)
			ctx = context.WithValue(ctx, "BucketKey", bucketKey)
			ctx = context.WithValue(ctx, "AWSRegion", awsRegion)
			ctx = context.WithValue(ctx, "TableThree", tableThree)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Endpoint Hit: /")
		fmt.Fprintf(w, "Welcome to Eli's Backend!")
	})

	eliHandler := handlers.NewHandler(eliClient)
	router.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Endpoint Hit: /info")
		eliHandler.GetBasicInfo(w, r)
	}).Methods("GET")

	router.HandleFunc("/education", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Endpoint Hit: /education")
		eliHandler.GetEducationHistory(w, r)
	}).Methods("GET")

	router.HandleFunc("/summary", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Endpoint Hit: /summary")
		eliHandler.GetSummary(w, r)
	}).Methods("GET")

	router.HandleFunc("/experience", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Endpoint Hit: /experience")
		eliHandler.GetExperienceHistory(w, r)
	}).Methods("GET")

	router.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Endpoint Hit: /resume")
		eliHandler.GetResume(w, r)
	}).Methods("GET")

	router.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Endpoint Hit: /projects")
		eliHandler.GetProjects(w, r)
	}).Methods("GET")

	handler := cors.SetCORSHeader(router)

	port := 8000
	fmt.Printf("Server is running on :%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	fmt.Println("Application started successfully")
}
