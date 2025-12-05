package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/kolesico/FocusGuard/internal/application"
)

func main() {
	configPath := flag.String("config-path", "./config/config.yaml", "path to config file")

	flag.Parse()

	app, err := application.NewApp(*configPath)

	if err != nil {
		log.Fatal("Error initionalize application")
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}

}

// runTestServer Запуск тестового слушающего сервера
func runTestServer() {
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello, Go!")
	})

	fmt.Println("Start server 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error start server")
	}
}