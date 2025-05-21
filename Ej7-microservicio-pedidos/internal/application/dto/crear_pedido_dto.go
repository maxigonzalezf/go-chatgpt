// DTOs de entrada/salida
// Pensado para contener los Data Transfer Objects especificos del caso de uso CrearPedido
// Es una estructura que usamos para:
//   - Recibir datos desde una capa externa (handler HTTP, CLI, etc)
//   - Devolver datos hacia esa capa, sin exponer directamente entidades del dominio
package dto

import "github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"

type CrearPedidoInput struct {
	Cliente string
	Monto float64
	Moneda string
}

type CrearPedidoOutput struct {
	ID string
	Total domain.Dinero
}