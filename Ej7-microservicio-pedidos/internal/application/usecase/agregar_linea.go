// internal/application/usecase/agregar_linea.go
package usecase

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/application/dto"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"
)

type AgregarLineaUseCase struct {
	Repo domain.LineaPedidoRepository
}

func (uc *AgregarLineaUseCase) Ejecutar(in dto.AgregarLineaInput) error {
	if in.Cantidad <= 0 {
		return fmt.Errorf("la cantidad debe ser mayor a cero")
	}

	linea := domain.LineaPedido{
		ID:       uuid.New().String(),
		PedidoID: in.PedidoID,
		Producto: in.Producto,
		Cantidad: in.Cantidad,
		Precio:   in.Precio,
	}
	log.Printf("Buscando pedido con ID: %s", in.PedidoID)

	return uc.Repo.SaveLinea(linea)
}
