// Puerto de salida: interfaz PedidoRepository
package domain

type PedidoRepository interface {
	Save(p Pedido) error
	FindByID(id string) (Pedido, error)
}