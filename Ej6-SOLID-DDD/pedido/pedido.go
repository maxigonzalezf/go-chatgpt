package pedido

import (
	"errors"

	"github.com/maxigonzalezf/go-chatgpt/Ej6-SOLID-DDD/dinero"
)

type Pedido struct {
	ID    string
	Total dinero.Dinero
}

// Servicio de dominio: agregar una línea de monto
func AgregarLinea(p *Pedido, monto dinero.Dinero) error {
	if p.Total.Moneda != "" && !p.Total.EsMismaMoneda(monto) {
		return errors.New("la moneda del monto no coincide con la del pedido")
	}

	// Si es la primera línea, se inicializa Total con la moneda
	if p.Total.Moneda == "" {
		p.Total = monto
		return nil
	}

	total, err := p.Total.Sumar(monto)
	if err != nil {
		return err
	}

	p.Total = total
	return nil
}
