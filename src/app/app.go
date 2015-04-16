package main

import (
    "encoding/json"
    "net/http"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func main() {
	api := web.New()
    goji.Handle("/api/*", api)

    api.Use(middleware.SubRouter)

    api.Get("/tasks", ListTasks)

    goji.Serve()
}

var tasks = []Task{
	Task{ID: 1, Title: "spam", Description: "abc", DueDate: "2015/10/10", Done: false},
	Task{ID: 2, Title: "egg", Description: "efg", DueDate: "2015/06/10",Done: false},
}

type Task struct {
	ID			uint16
    Title       string
    Description string
    DueDate		string
    Done        bool
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    encoder := json.NewEncoder(w)
    encoder.Encode(tasks)
}

