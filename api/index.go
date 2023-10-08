package handler

import (
	"example/go-taco/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func Handler(w http.ResponseWriter, request *http.Request) {
	r := chi.NewRouter()
	r.Get("/api/taco", handlers.List)
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("Erro ao subir o server:", err)
		return
	}
}
