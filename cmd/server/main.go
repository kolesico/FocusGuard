package main


import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Go!")
	})

	fmt.Println("Start server 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error start server")
	}
}