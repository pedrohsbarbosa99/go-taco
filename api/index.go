package main

import (
	"example/go-taco/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/api/taco", handlers.List)
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("Erro ao subir o server:", err)
		return
	}

}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}