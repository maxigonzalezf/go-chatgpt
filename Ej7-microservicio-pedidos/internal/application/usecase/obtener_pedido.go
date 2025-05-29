package usecase

import (
	"fmt"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/dto"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"
)

type ObtenerPedidoUseCase struct {
	Repo domain.PedidoRepository
}

func (uc *ObtenerPedidoUseCase) Ejecutar(id string) (dto.CrearPedidoOutput, error) {
	pedido, err := uc.Repo.FindByID(id)
	if err != nil {
		return dto.CrearPedidoOutput{}, fmt.Errorf("pedido no encontrado")
	}

	out := dto.CrearPedidoOutput{
		ID:    pedido.ID,
		Total: pedido.Total,
	}
	return out, nil
}
