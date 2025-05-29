// internal/domain/linea_pedido_repository.go
package domain

type LineaPedidoRepository interface {
	SaveLinea(LineaPedido) error
	FindLineByPedidoID(id string) ([]LineaPedido, error)
}
