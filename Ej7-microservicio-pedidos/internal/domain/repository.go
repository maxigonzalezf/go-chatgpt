// Puerto de salida: interfaz PedidoRepository
package domain

// Define que operaciones necesita el dominio, no como se implementan

type PedidoRepository interface {
	Save(p Pedido) error
	FindByID(id string) (Pedido, error)
}

// Puertos de salida: abstraen la persistencia
// Polimorfismo estatico: Go asigna implementaciones que satisfacen la interfaz