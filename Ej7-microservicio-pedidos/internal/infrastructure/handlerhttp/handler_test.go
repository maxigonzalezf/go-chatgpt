package handlerhttp

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/dto"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/usecase"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/infrastructure/persistence"
)

var testDSN = "host=localhost port=5432 user=postgres password=secret dbname=pedidos_test sslmode=disable"

// Funcion auxiliar para tests.
// Se conecta a la BDD de produccion, devuelve una implementacion real del repo para testear la app de manera controlada
// Limpia la bdd antes de cada test
func setupSQLRepo(t *testing.T) *persistence.PedidoRepoSQL {
	db, err := sql.Open("postgres", testDSN)
	if err != nil {
		t.Fatalf("no pude abrir BD: %v", err)
	}
	// Limpieza previa
	if _, err := db.Exec("TRUNCATE pedidos"); err != nil {
		t.Fatalf("error limpiando la tabla pedidos: %v", err)
	}
	return persistence.NewPedidoRepoSQL(db)
}

// 201, JSON con ID y Total
func TestCrearPedidoHandler_Success(t *testing.T) {
	// 1. Prepara repo y usecase
	// repo := persistence.NewPedidoRepoMemoria() // repo en memoria
	repo := setupSQLRepo(t) // repo DB SQL
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

// Invalid JSON, 400, mensaje de formato
func TestCrearPedidoHandler_BadRequest_InvalidJSON(t *testing.T) {
	// repo := persistence.NewPedidoRepoMemoria() // repo en memoria
	repo := setupSQLRepo(t) // repo DB SQL
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

// Business error, monto negativo -> 400, mensaje de dominio
func TestCrearPedidoHandler_BadRequest_LogicalError(t *testing.T) {
	// repo := persistence.NewPedidoRepoMemoria() // repo en memoria
	repo := setupSQLRepo(t) // repo DB SQL
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

// Table-driven tests para multiples escenarios
// Capturar y verificar stdout (en su caso)
// Separacion de pruebas: tests puramente del dominio vs tests de integracion ligera HTTP

func TestObtenerPedidoHandler_SQL_Success(t *testing.T) {
	repo := setupSQLRepo(t)
	obtenerUC := usecase.ObtenerPedidoUseCase{Repo: repo}

	// Guardamos un pedido en la base de datos
	id := "sql123"
	pedido := domain.Pedido{
		ID: id,
		Total: domain.Dinero{
			Moneda:   "USD",
			Cantidad: 200,
		},
	}
	if err := repo.Save(pedido); err != nil {
		t.Fatalf("error guardando pedido en SQL: %v", err)
	}

	// Creamos un request a /pedidos/sql123
	req := httptest.NewRequest(http.MethodGet, "/pedidos/"+id, nil)
	rr := httptest.NewRecorder()

	// Usamos el subrouter como handler
	handler := PedidosSubrouter(&obtenerUC, nil, nil)
	handler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d; want %d", rr.Code, http.StatusOK)
	}

	var out dto.CrearPedidoOutput
	if err := json.NewDecoder(rr.Body).Decode(&out); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}
	if out.ID != id || out.Total.Moneda != "USD" || out.Total.Cantidad != 200 {
		t.Errorf("respuesta inesperada: %+v", out)
	}
}


func TestObtenerPedidoHandler_NotFound(t *testing.T) {
	repo := setupSQLRepo(t)
	obtenerUC := usecase.ObtenerPedidoUseCase{Repo: repo}
	handler := PedidosSubrouter(&obtenerUC, nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/pedidos/no-existe", nil)
	rr := httptest.NewRecorder()

	handler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("status = %d; want %d", rr.Code, http.StatusNotFound)
	}

	if !strings.Contains(rr.Body.String(), "no encontrado") {
		t.Errorf("mensaje = %q; esperaba error", rr.Body.String())
	}
}

