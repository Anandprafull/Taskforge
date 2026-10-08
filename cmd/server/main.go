package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"taskforge/internal/database"
	"taskforge/internal/jobs"
)

var jobID = 1
var jobList = []jobs.Job{}

func createJob(w http.ResponseWriter, r *http.Request) {
	var job jobs.Job

	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	job.ID = jobID
	jobID++

	jobList = append(jobList, job)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

func getJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobList)
}


func main() {

	conn,err:=database.Connect()
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	fmt.Println("Connected to PostgresSQL!")

	err = database.CreateTables(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database ready!")


	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "TaskForge API is running!")
	})

	http.HandleFunc("POST /jobs", createJob)

	http.HandleFunc("GET /jobs",getJobs)

	fmt.Println("Server running on :8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}