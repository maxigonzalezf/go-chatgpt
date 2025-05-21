// Punto de entrada: arranca el HTTP server
package main

import (
	"log"
	"net/http"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/usecase"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/infrastructure/persistence"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/infrastructure/handlerhttp"

)

func main() {
	repo := persistence.NewPedidoRepoMemoria()
	uc := usecase.CrearPedidoUseCase{Repo: repo}
	mux := http.NewServeMux()
	mux.HandleFunc("/pedidos", handlerhttp.CrearPedidoHandler(&uc))

	log.Println("Servidor iniciado en :8080")
	http.ListenAndServe(":8080", mux)
}