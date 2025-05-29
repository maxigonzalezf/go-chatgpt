// internal/domain/linea_pedido.go
package domain

type LineaPedido struct {
	ID        string
	PedidoID  string
	Producto  string
	Cantidad  int
	Precio    Dinero // precio unitario
}
