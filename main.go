package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/romen95/go_final_project/app/controller"
	"github.com/romen95/go_final_project/app/database"

	"github.com/go-chi/chi/v5"
)

const PORT = 7540

func getPort() int {
	p, _ := strconv.ParseInt(os.Getenv("TODO_PORT"), 10, 32)
	if p == 0 {
		return PORT
	}
	return int(p)
}

func main() {
	database.InstallDb()
	webDir := "./web"

	r := chi.NewRouter()
	r.Mount("/", http.FileServer(http.Dir(webDir)))
	r.Get("/api/nextdate", controller.NextDate)
	r.Post("/api/task", controller.AddTask)
	r.Get("/api/tasks", controller.TasksReadGET)
	r.Get("/api/task", controller.TaskReadGET)
	r.Put("/api/task", controller.TaskUpdatePUT)
	r.Post("/api/task/done", controller.TaskDonePOST)
	r.Delete("/api/task", controller.TaskDELETE)

	serverPort := getPort()
	log.Println(fmt.Sprintf("Server started on port: %d", serverPort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), r))
}
