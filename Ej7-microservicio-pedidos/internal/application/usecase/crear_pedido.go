// Caso de uso: CrearPedidoUseCase
package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/dto"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"
)

type CrearPedidoUseCase struct {
	Repo domain.PedidoRepository
	PedidosChan chan<- string // canal inyectado desde main
}

func (uc *CrearPedidoUseCase) Ejecutar(in dto.CrearPedidoInput) (dto.CrearPedidoOutput, error) {
	// 1. Validar input
	if in.Monto < 0 {
		return dto.CrearPedidoOutput{}, fmt.Errorf("debe ingresar un monto correcto")
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
	// 4. Repo.Save(pedido) -> persistencia a traves del puerto de salida
	if err := uc.Repo.Save(p); err != nil {
		return dto.CrearPedidoOutput{}, err
	}
	
	// 5. Enviar el ID al canal para procesamiento en background
	uc.PedidosChan <- p.ID

	// 6. Construir y devolver output (creacion DTO de salida)
	out := dto.CrearPedidoOutput {
		ID: p.ID,
		Total: p.Total,
	}

	return out, nil
}

// Separacion de responsabilidades: la logica de negocio esta aca, aislada de HTTP y de la infraestructura
// Dependency Inversion: el caso de uso solo conoce la abstraccion PedidoRepository
// Value Object: Dinero modela cantidad y moneda, con reglas propias (si se agrega Sumar, EsMismaMoneda, etc)