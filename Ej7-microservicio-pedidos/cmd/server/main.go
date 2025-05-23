// Punto de entrada: arranca el HTTP server
package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/usecase"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/infrastructure/handlerhttp"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/infrastructure/persistence"
)

// El main orquesta, pero no contiene logica de negocio
// Solo se configura el HTTP server y se enlaza con el caso de uso

func main() {
	// se importa el repo en memoria
	//repo := persistence.NewPedidoRepoMemoria()
	// se construye el caso de uso inyectandole el repo
	//uc := usecase.CrearPedidoUseCase{Repo: repo}
	//mux := http.NewServeMux()
	// se registra la ruta /pedidos apuntando a CrearPedidoHandler
	//mux.HandleFunc("/pedidos", handlerhttp.CrearPedidoHandler(&uc))

	//log.Println("Servidor iniciado en :8080")
	//http.ListenAndServe(":8080", mux)

	// 1. Conexión a la base de datos
	dsn := "postgres://postgres:secret@localhost:5432/pedidos_db?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error abriendo BD: %v", err)
	}
	defer db.Close()

	// 2. Ping para asegurarnos de la conexión
	if err := db.Ping(); err != nil {
		log.Fatalf("error conectando a BD: %v", err)
	}

	// 3. Crear adaptador SQL en lugar de memoria
	repoSQL := persistence.NewPedidoRepoSQL(db)

	// 4. Inyectar en el caso de uso
	uc := usecase.CrearPedidoUseCase{Repo: repoSQL}

	// 5. Montar handler y servidor HTTP
	mux := http.NewServeMux()
	mux.HandleFunc("/pedidos", handlerhttp.CrearPedidoHandler(&uc))

	log.Println("Servidor escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
