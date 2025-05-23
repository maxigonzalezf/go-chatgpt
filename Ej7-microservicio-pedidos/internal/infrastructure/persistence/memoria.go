// Adaptador de salida: repo en memoria
package persistence

import (
	"errors"
	"sync"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"
)

type PedidoRepoMemoria struct {
	mu sync.Mutex // para la concurrencia
	datos map[string]domain.Pedido // para guardar pedidos
}

func NewPedidoRepoMemoria() *PedidoRepoMemoria {
	return &PedidoRepoMemoria{datos: make(map[string]domain.Pedido)}
}

// Implementacion de los metodos de la interface (guardan y recuperan desde el map)

func (r *PedidoRepoMemoria) Save(p domain.Pedido) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.datos[p.ID] = p

	return nil
}

func (r *PedidoRepoMemoria) FindByID(id string) (domain.Pedido, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, ok := r.datos[id]
	if !ok {
		return domain.Pedido{}, errors.New("no encontrado")
	}

	return p, nil
}

// Adaptadores: un adaptador concreto que satisface el puerto
// Test doubles: este repo se usa tanto en produccion como en test de integracion