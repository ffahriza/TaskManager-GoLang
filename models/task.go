package models

type Task struct {
    ID    string `json:"id" bson:"id"`
    Title string `json:"title" bson:"title"`
    Done  bool   `json:"done" bson:"done"`
}