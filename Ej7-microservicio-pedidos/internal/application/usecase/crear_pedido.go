// Caso de uso: CrearPedidoUseCase
package usecase

import (
	"fmt"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"
	"github.com/google/uuid"
)

type CrearPedidoInput struct {
	Cliente string
	Monto float64
	Moneda string
}

type CrearPedidoOutput struct {
	ID string
	Total domain.Dinero
}

type CrearPedidoUseCase struct {
	Repo domain.PedidoRepository
}

func (uc *CrearPedidoUseCase) Ejecutar(in CrearPedidoInput) (CrearPedidoOutput, error) {
	// 1. Validar input
	if in.Monto < 0 {
		return CrearPedidoOutput{}, fmt.Errorf("debe ingresar un monto correcto")
	}
	// 2. Construir domain.Dinero
	total := domain.Dinero {
		Moneda: in.Moneda,
		Cantidad: in.Monto,
	}
	// 3. Crear la entidad Pedido (con ID generado (uuid))
	id := uuid.New().String()
	p := domain.Pedido{
		ID: id,
		Total: total,
	}
	// 4. Repo.Save(pedido)
	if err := uc.Repo.Save(p); err != nil {
		return CrearPedidoOutput{}, err
	}
	// 5. Devolver output (DTO de salida)
	out := CrearPedidoOutput {
		ID: p.ID,
		Total: p.Total,
	}
	return out, nil
}