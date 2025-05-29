package usecase

import (
	"testing"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/dto"
)

func TestObtenerPedidoUseCase_Success(t *testing.T) {
    repo := &RepoFake{}
    
    // Creamos un pedido con el usecase
    crearUC := CrearPedidoUseCase{Repo: repo}
    input := dto.CrearPedidoInput{
        Cliente: "Maxi",
        Monto:   100.0,
        Moneda:  "ARS",
    }
    creado, err := crearUC.Ejecutar(input)
    if err != nil {
        t.Fatalf("no se pudo crear pedido de prueba: %v", err)
    }

    // Lo buscamos con el usecase de obtener
    obtenerUC := ObtenerPedidoUseCase{Repo: repo}
    out, err := obtenerUC.Ejecutar(creado.ID)
    if err != nil {
        t.Fatalf("error inesperado: %v", err)
    }

    if out.ID != creado.ID || out.Total.Moneda != input.Moneda || out.Total.Cantidad != input.Monto {
        t.Errorf("salida inesperada: %+v", out)
    }
}

