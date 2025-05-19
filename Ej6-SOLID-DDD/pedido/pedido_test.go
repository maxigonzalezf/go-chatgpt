package pedido

import (
	"testing"

	"github.com/maxigonzalezf/go-chatgpt/Ej6-SOLID-DDD/dinero"
)

func TestAgregarLinea(t *testing.T) {
	// Que se pueda iniciar un pedido sin moneda y se asigne la del primer monto
	t.Run("Agregar primera linea (inicializa moneda)", func(t *testing.T) {
		p := Pedido{ID: "P-001"}
		m := dinero.Dinero{Moneda: "USD", Cantidad: 100}

		err := AgregarLinea(&p, m)

		if err != nil {
			t.Errorf("No esperaba error, pero obtuve: %v", err)
		}
		if p.Total.Moneda != "USD" || p.Total.Cantidad != 100 {
			t.Errorf("Total esperado USD 100, pero obtuve %v %v", p.Total.Moneda, p.Total.Cantidad)
		}
	})

	// Que el total se sume correctamente sin errores
	t.Run("Agregar línea válida con misma moneda", func(t *testing.T) {
		p := Pedido{
			ID:    "P-002",
			Total: dinero.Dinero{Moneda: "USD", Cantidad: 50},
		}
		m := dinero.Dinero{Moneda: "USD", Cantidad: 25}

		err := AgregarLinea(&p, m)

		if err != nil {
			t.Errorf("No esperaba error, pero obtuve: %v", err)
		}
		if p.Total.Cantidad != 75 {
			t.Errorf("Esperaba total de 75, pero obtuve %v", p.Total.Cantidad)
		}
	})

	// Que se dispare un error segun la regla de negocio
	t.Run("Falla por moneda distinta", func(t *testing.T) {
		p := Pedido{
			ID:    "P-003",
			Total: dinero.Dinero{Moneda: "USD", Cantidad: 30},
		}
		m := dinero.Dinero{Moneda: "EUR", Cantidad: 20}

		err := AgregarLinea(&p, m)

		if err == nil {
			t.Errorf("Esperaba error por moneda distinta, pero no lo hubo")
		}
	})
}
