package tasks

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	Decription string `json:"tasks"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	task := Task{Decription: "Some description"}
	json.NewEncoder(w).Encode(task)

}
