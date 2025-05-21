// Adaptador de entrada: HTTP handler
package handlerhttp

import (
	"encoding/json"
	"net/http"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/usecase"
)

func CrearPedidoHandler(uc *usecase.CrearPedidoUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input usecase.CrearPedidoInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "formato de JSON inválido", http.StatusBadRequest)
			return
		}
		output, err := uc.Ejecutar(input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(output)
	}
}
