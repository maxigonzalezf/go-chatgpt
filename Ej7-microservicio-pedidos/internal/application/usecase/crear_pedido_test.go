package usecase

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/dto"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"
)

// Mock de repository
type RepoFake struct {
	Saved domain.Pedido
	Err error
	datos map[string]domain.Pedido
}

func (r *RepoFake) Save(p domain.Pedido) error {
	r.Saved = p
	if r.datos == nil {
		r.datos = make(map[string]domain.Pedido)
	}
	r.datos[p.ID] = p
	return r.Err
}

func (r *RepoFake) FindByID(id string) (domain.Pedido, error) {
    p, ok := r.datos[id]
    if !ok {
        return domain.Pedido{}, fmt.Errorf("pedido no encontrado")
    }
    return p, nil
}

func TestCrearPedido_Ejecutar(t *testing.T) {
	repo := &RepoFake{}
	uc := CrearPedidoUseCase{Repo: repo}

	in := dto.CrearPedidoInput{Cliente: "Mara", Monto: 123.45, Moneda: "USD"}
	out, err := uc.Ejecutar(in)
	if err != nil {
		t.Fatalf("no esperaba error, pero lo tuve: %v", err)
	}
	if out.ID == "" {
		t.Error("esperaba un ID generado, vino vacio")
	}
	if out.Total.Cantidad != in.Monto || out.Total.Moneda != in.Moneda {
		t.Errorf("total mal: got %v, want %v %v", out.Total, in.Monto, in.Moneda)
	}
	// Verifica que el repo recibio el mismo pedido
	if repo.Saved.ID != out.ID {
		t.Errorf("pedido generado con ID distinto: got %q, want %q", repo.Saved.ID, out.ID)
	}
}

func TestCrearPedido_MontoNegativo(t *testing.T) {
	uc := CrearPedidoUseCase{Repo: &RepoFake{}}
	_, err := uc.Ejecutar(dto.CrearPedidoInput{Cliente: "Mara", Monto:  -5, Moneda: "USD"})
	if err == nil || !strings.Contains(err.Error(), "debe ingresar un monto correcto") {
		t.Errorf("esperaba error de monto invalido, pero obtuve: %v", err)
	}
}