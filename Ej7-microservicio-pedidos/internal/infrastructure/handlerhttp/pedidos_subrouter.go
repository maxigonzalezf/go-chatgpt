package handlerhttp

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/dto"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/usecase"
)

func PedidosSubrouter(
	obtenerUC *usecase.ObtenerPedidoUseCase,
	agregarLineaUC *usecase.AgregarLineaUseCase,
	obtenerLineasUC *usecase.ObtenerLineasUseCase,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/pedidos/")
		parts := strings.Split(path, "/")

		if len(parts) < 1 || parts[0] == "" {
			http.Error(w, "ID de pedido no especificado", http.StatusBadRequest)
			return
		}
		pedidoID := parts[0]

		if len(parts) == 1 {
			if r.Method == http.MethodGet {
				obtenerPedidoHandler(obtenerUC, pedidoID, w, r)
				return
			}
			http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
			return
		}

		if len(parts) == 2 && parts[1] == "lineas" {
			switch r.Method {
			case http.MethodPost:
				agregarLineaHandler(agregarLineaUC, pedidoID, w, r)
			case http.MethodGet:
				obtenerLineasHandler(obtenerLineasUC, pedidoID, w, r)
			default:
				http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
			}
			return
		}

		http.Error(w, "ruta no válida", http.StatusNotFound)
	}
}

func obtenerPedidoHandler(uc *usecase.ObtenerPedidoUseCase, id string, w http.ResponseWriter, r *http.Request) {
	out, err := uc.Ejecutar(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}

func agregarLineaHandler(uc *usecase.AgregarLineaUseCase, pedidoID string, w http.ResponseWriter, r *http.Request) {
	var input dto.AgregarLineaInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "input inválido", http.StatusBadRequest)
		return
	}
	input.PedidoID = pedidoID
	if err := uc.Ejecutar(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func obtenerLineasHandler(uc *usecase.ObtenerLineasUseCase, pedidoID string, w http.ResponseWriter, r *http.Request) {
	lineas, err := uc.Ejecutar(pedidoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lineas)
}
