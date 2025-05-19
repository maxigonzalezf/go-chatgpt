package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// PagoFake simula un MedioDePago y registra la llamada.
type PagoFake struct {
	LlamadoCon float64
	OKReturn   bool
}

// COMPLETA ESTE MÉTODO:
// Pagar debe guardar el monto en p.LlamadoCon y devolver p.OKReturn
func (p *PagoFake) Pagar(monto float64) bool {
	// TODO: implementar
	p.LlamadoCon = monto
	return p.OKReturn
}

// Test que comprueba que ProcesarPago llama a Pagar con el monto correcto.
func TestProcesarPago_LlamaPagar(t *testing.T) {
	fake := &PagoFake{OKReturn: true}
	ProcesarPago(fake, 123.45)

	if fake.LlamadoCon != 123.45 {
		t.Errorf("Pagar fue llamado con %f, quiero %f", fake.LlamadoCon, 123.45)
	}
}

// Test con tabla de escenarios para OKReturn = true/false
func TestProcesarPago_VariosEscenarios(t *testing.T) {
	casos := []struct {
		name     string
		okReturn bool
		monto    float64
	}{
		{"Éxito", true, 50},
		{"Rechazado", false, 75},
	}

	for _, c := range casos {
		t.Run(c.name, func(t *testing.T) {
			fake := &PagoFake{OKReturn: c.okReturn}
			// Capturamos stdout
			original := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			
			// 1) Llamar a ProcesarPago con fake y c.monto
			ProcesarPago(fake, c.monto)

			// Restauramos stdout y leemos lo que se imprimió
			w.Close()
			os.Stdout = original
			out, _ := io.ReadAll(r)
			salida := string(out)
			// 2) Verificar que fake.LlamadoCon == c.monto
			if fake.LlamadoCon != c.monto {
				t.Errorf("Pagar fue llamado con %f, quiero %f", fake.LlamadoCon, 123.45)
			}
			// 3) (Opcional) podrías comprobar algo extra según c.okReturn
			// Verificamos que el mensaje correcto fue impreso
			if c.okReturn && !strings.Contains(salida, "éxito") {
				t.Errorf("Esperaba mensaje de éxito, pero se imprimió: %q", salida)
			}
			if !c.okReturn && !strings.Contains(salida, "rechazado") {
				t.Errorf("Esperaba mensaje de rechazo, pero se imprimió: %q", salida)
			}
		})
	}
}
