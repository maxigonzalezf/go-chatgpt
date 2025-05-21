package handlerhttp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/dto"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/usecase"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/infrastructure/persistence"
)

func TestCrearPedidoHandler_Success(t *testing.T) {
	// 1. Prepara repo y usecase
	repo := persistence.NewPedidoRepoMemoria()
	uc := usecase.CrearPedidoUseCase{Repo: repo}
	h := CrearPedidoHandler(&uc) // inyectamos el repo en CrearPedidoUseCase

	// 2. Crea el body JSON de la petición POST
	input := dto.CrearPedidoInput{
		Cliente: "Mara",
		Monto:   123.45,
		Moneda:  "USD",
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(input); err != nil {
		t.Fatalf("error encoding input: %v", err)
	}

	// 3. Construye la petición y el recorder
	req := httptest.NewRequest(http.MethodPost, "/pedidos", buf)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder() // intercepta lo que el handler escribiria a http.ResponseWriter

	// 4. Ejecuta el handler
	h(rr, req)

	// 5. Comprueba el status code
	if rr.Code != http.StatusCreated {
		t.Errorf("status = %d; want %d", rr.Code, http.StatusCreated)
	} // aseguramos que devuelva 201 Created

	// 6. Parsea el response body
	body, _ := io.ReadAll(rr.Body)
	var out dto.CrearPedidoOutput
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("error unmarshaling response: %v\nbody: %s", err, string(body))
	}

	// 7. Verifica contenido del DTO
	if out.Total.Moneda != input.Moneda || out.Total.Cantidad != input.Monto {
		t.Errorf("output.Total = %+v; want %+v", out.Total, input)
	}
	if out.ID == "" {
		t.Error("expected non-empty ID in response")
	}

	// 8. Verifica que el repo haya guardado el mismo pedido
	saved, err := repo.FindByID(out.ID)
	if err != nil {
		t.Fatalf("repo.FindByID error: %v", err)
	}
	if saved.ID != out.ID {
		t.Errorf("repo saved ID = %q; want %q", saved.ID, out.ID)
	}
}

/* Explicación paso a paso
Repositorio y UseCase
	Creamos un PedidoRepoMemoria real y lo inyectamos en CrearPedidoUseCase.

Request
	Serializamos un CrearPedidoInput a JSON y lo ponemos en el body de una petición POST.

Recorder
	httptest.NewRecorder() intercepta lo que tu handler escribiría a http.ResponseWriter.

Ejecución
	Llamas directamente al handler como función: h(rr, req).

Verificación HTTP
	Aseguramos que devuelva 201 Created.

Parseo JSON
	Leemos rr.Body, lo unmarshaleamos en CrearPedidoOutput y comprobamos campos.

Chequeo de persistencia
	Llamamos a repo.FindByID para confirmar que el pedido efectivamente se guardó.

Con esto cubrís un test de integración ligero que valida todo el flujo desde HTTP hasta dominio y repo en memoria.
*/

func TestCrearPedidoHandler_BadRequest_InvalidJSON(t *testing.T) {
	repo := persistence.NewPedidoRepoMemoria()
	uc := usecase.CrearPedidoUseCase{Repo: repo}
	h := CrearPedidoHandler(&uc)

	// 1. Payload que no es JSON valido
	badBody := strings.NewReader("{invalid-json") // body que no se puede decodificar
	req := httptest.NewRequest(http.MethodPost, "/pedidos", badBody)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	h(rr, req)

	// 2. Debe responder 400
	if rr.Code != http.StatusBadRequest {
		t.Errorf("status = %d; want %d", rr.Code, http.StatusBadRequest)
	}

	// 3. Mensaje de error generico en el body
	if !strings.Contains(rr.Body.String(), "formato") && !strings.Contains(rr.Body.String(), "error") {
		t.Errorf("body = %q; want error message", rr.Body.String())
	}
}

func TestCrearPedidoHandler_BadRequest_LogicalError(t *testing.T) {
	repo := persistence.NewPedidoRepoMemoria()
	uc := usecase.CrearPedidoUseCase{Repo: repo}
	h := CrearPedidoHandler(&uc)

	// 1. Payload valido como JSON pero con monto no permitido (<0)
	payload := dto.CrearPedidoInput{
		Cliente: "Test", Monto: -10, Moneda: "USD",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(payload)

	req := httptest.NewRequest(http.MethodPost, "/pedidos", buf)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	h(rr, req)

	// 2. Debe responder 400
	if rr.Code != http.StatusBadRequest {
		t.Errorf("status = %d; want %d", rr.Code, http.StatusBadRequest)
	}
	
	// 3. Contenido del error debe mencionar la regla de negocio
	if !strings.Contains(rr.Body.String(), "monto") {
		t.Errorf("body = %q; want mention of monto error", rr.Body.String())
	}
}
