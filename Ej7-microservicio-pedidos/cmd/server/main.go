// Punto de entrada: arranca el HTTP server
package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/usecase"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/infrastructure/handlerhttp"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/infrastructure/persistence"
)

var pedidosChan = make(chan string)

func iniciarWorker() {
    go func() {
        for id := range pedidosChan {
            log.Printf("Worker: procesando pedido %s", id)
            // Podés simular enviar un email u otra tarea
            time.Sleep(1 * time.Second)
            log.Printf("Worker: finalizó pedido %s", id)
        }
    }()
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Inicio: %s %s", r.Method, r.URL.Path)

        next.ServeHTTP(w, r)

        log.Printf("Fin: %s %s (duración: %s)", r.Method, r.URL.Path, time.Since(start))
    })
}


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

	// AGREGAMOS UN PEDIDO HARDCODEADO PARA PRUEBA
	pedido := domain.Pedido{
		ID: "abc123",
		Total: domain.Dinero{
			Moneda:   "USD",
			Cantidad: 99.99,
		},
	}
	if err := repoSQL.Save(pedido); err != nil {
		log.Fatal("no se pudo guardar pedido de prueba:", err)
	}

	// 4. Inyectar en el caso de uso
	crearUc := usecase.CrearPedidoUseCase{Repo: repoSQL, 
										  PedidosChan: pedidosChan,
										}
	obtenerUc := usecase.ObtenerPedidoUseCase{Repo: repoSQL}
	agregarLineaUc := usecase.AgregarLineaUseCase{Repo: repoSQL}
	obtenerLineasUc := usecase.ObtenerLineasUseCase{Repo: repoSQL}

	// 5. Montar handler y servidor HTTP
	mux := http.NewServeMux()
	crearHandler := http.HandlerFunc(handlerhttp.CrearPedidoHandler(&crearUc))
	mux.Handle("/pedidos", LoggingMiddleware(crearHandler))
	//mux.HandleFunc("/pedidos", handlerhttp.CrearPedidoHandler(&crearUc))
	mux.HandleFunc("/pedidos/", handlerhttp.PedidosSubrouter(
	&obtenerUc,
	&agregarLineaUc,
	&obtenerLineasUc,
))

	iniciarWorker()
	log.Println("Servidor escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
