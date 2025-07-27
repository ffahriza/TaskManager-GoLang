package handlers

import (
    "encoding/json"
    "net/http"
    "context"
    "github.com/ffahriza/TaskManager-GoLang/models"
    "github.com/ffahriza/TaskManager-GoLang/db"

)

var tasks = []models.Task{
    {ID: "1", Title: "Learn Go", Done: false},
    {ID: "2", Title: "Build API", Done: false},
}

// List all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    _ = json.NewDecoder(r.Body).Decode(&task)

    _, err := db.TaskCollection.InsertOne(context.TODO(), task)
    if err != nil {
        http.Error(w, "Failed to create task", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "UpdateTask not implemented", http.StatusNotImplemented)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "DeleteTask not implemented", http.StatusNotImplemented)
}