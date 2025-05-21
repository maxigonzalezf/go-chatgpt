// Adaptador de salida: repo en memoria
package persistence

import (
	"errors"
	"sync"

	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"
)

type PedidoRepoMemoria struct {
	mu sync.Mutex
	datos map[string]domain.Pedido
}

func NewPedidoRepoMemoria() *PedidoRepoMemoria {
	return &PedidoRepoMemoria{datos: make(map[string]domain.Pedido)}
}

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