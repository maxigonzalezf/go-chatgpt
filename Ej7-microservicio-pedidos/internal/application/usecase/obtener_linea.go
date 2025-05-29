package usecase

import "github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"

type ObtenerLineasUseCase struct {
	Repo domain.LineaPedidoRepository
}

func (uc *ObtenerLineasUseCase) Ejecutar(pedidoID string) ([]domain.LineaPedido, error) {
	return uc.Repo.FindLineByPedidoID(pedidoID)
}
