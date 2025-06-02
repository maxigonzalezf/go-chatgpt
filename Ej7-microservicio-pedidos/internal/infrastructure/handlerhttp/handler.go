// Adaptador de entrada: HTTP handler
package handlerhttp

import (
	"encoding/json"
	"net/http"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/dto"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/usecase"
)

func CrearPedidoHandler(uc *usecase.CrearPedidoUseCase) http.HandlerFunc {
	// El handler recibe un http.Request con JSON en el body
	// Se decodifica y valida el JSON y se lo convierte al DTO CrearPedidoInput
	return func(w http.ResponseWriter, r *http.Request) {
		var input dto.CrearPedidoInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "formato de JSON inválido", http.StatusBadRequest)
			return
		}
		// Se invoca el caso de uso pasandole el DTO (si hay error, devuelve 400 BadRequest)
		output, err := uc.Ejecutar(input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// si no hay error, 201 Created y se codifica al DTO CrearPedidoOutput
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(output)
	}
}

// Puertos de entrada: la firma http.HandlerFunc es el adaptador que "empuja" la peticion al nucleo
// DTOs separan el formato HTTP del dominio
// Errores se traducen a codigos HTTP adecuados
