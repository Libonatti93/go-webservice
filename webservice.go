package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)

	// Define a porta em que o servidor irá rodar
	port := ":8080"
	fmt.Printf("Servidor rodando na porta %s\n", port)

	// Inicia o servidor
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
