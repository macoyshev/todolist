package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/macoyshev/todolist/internal/tasks"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/gTasks", tasks.GetTasks)
	server := http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	server.ListenAndServe()
}
