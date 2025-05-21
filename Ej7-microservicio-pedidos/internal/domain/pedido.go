// Entidad Pedido y Value Object Dinero
package domain

type Dinero struct {
	Moneda string
	Cantidad float64
}

type Pedido struct {
	ID string
	Total Dinero
}