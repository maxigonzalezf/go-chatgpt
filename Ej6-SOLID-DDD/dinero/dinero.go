package dinero

import "fmt"

type Dinero struct {
	Moneda  string
	Cantidad float64
}

func (d Dinero) EsMismaMoneda(otro Dinero) bool {
	return d.Moneda == otro.Moneda
}

func (d Dinero) Sumar(otro Dinero) (Dinero, error) {
	if !d.EsMismaMoneda(otro) {
		return Dinero{}, fmt.Errorf("no se puede sumar dinero en distintas monedas: %s y %s", d.Moneda, otro.Moneda)
	}
	
	return Dinero{
		Moneda:  d.Moneda,
		Cantidad: d.Cantidad + otro.Cantidad,
	}, nil
}
