package handlers

import (
    "encoding/json"
    "net/http"
    "context"
    "github.com/ffahriza/TaskManager-GoLang/models"
    "github.com/ffahriza/TaskManager-GoLang/db"
)

// POST /mongo/tasks
func CreateMongoTask(w http.ResponseWriter, r *http.Request) {
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