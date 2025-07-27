package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/ffahriza/TaskManager-GoLang/handlers"
    "github.com/ffahriza/TaskManager-GoLang/db"
)

// main function to start the server
func main() {
    // Initialize MongoDB connection
    db.InitMongo()

    // Create a new router
    r := mux.NewRouter()
    // Define routes
    r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
    r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

    // Start the server
    http.ListenAndServe(":8000", r)
}