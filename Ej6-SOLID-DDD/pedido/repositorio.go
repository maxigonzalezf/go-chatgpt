package pedido

type PedidoRepository interface {
	Save(p Pedido) error
	FindByID(id int) (Pedido, error)
}
