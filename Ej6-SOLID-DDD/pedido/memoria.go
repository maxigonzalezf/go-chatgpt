package pedido

import (
	"errors"
	"sync"
)

type RepositorioMemoria struct {
	mu       sync.Mutex
	datos    map[string]Pedido
}

func NuevoRepositorioMemoria() *RepositorioMemoria {
	return &RepositorioMemoria{
		datos: make(map[string]Pedido),
	}
}

func (r *RepositorioMemoria) Save(p Pedido) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.datos[p.ID] = p
	return nil
}

func (r *RepositorioMemoria) FindByID(id string) (Pedido, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, ok := r.datos[id]
	if !ok {
		return Pedido{}, errors.New("pedido no encontrado")
	}
	return p, nil
}
