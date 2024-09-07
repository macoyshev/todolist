package main

import (
	"net/http"

	"github.com/macoyshev/todolist/internal/tasks"
)

func main() {
	srv := http.NewServeMux()

	srv.HandleFunc("/gTasks", tasks.gTasks)
	srv.ListenAndServe()
}
