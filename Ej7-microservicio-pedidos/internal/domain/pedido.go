// Entidad Pedido y Value Object Dinero
package domain

type Dinero struct {
	Moneda string
	Cantidad float64
} // se le podrian sumar metodos como Sumar, EsMismaMoneda

type Pedido struct {
	ID string
	Total Dinero
} // se le pueden agregar metodos de dominio, p ej. AgregarLinea, CambiarEstado, etc para logica mas compleja

// Inmutabilidad parcial: idealmente el V.O no cambia, y la entidad controla su estado
// Reglas de negocio dentro de los metodos de dominio o de los casos de uso