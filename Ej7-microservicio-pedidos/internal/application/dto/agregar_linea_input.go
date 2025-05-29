// internal/application/dto/agregar_linea_input.go
package dto

import "github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"

type AgregarLineaInput struct {
	PedidoID string  `json:"-"`
	Producto string  `json:"producto"`
	Cantidad int     `json:"cantidad"`
	Precio   domain.Dinero  `json:"precio"`
}
