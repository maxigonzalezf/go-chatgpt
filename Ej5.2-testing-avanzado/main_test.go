package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// Mock de StockChecker
type StockFake struct {
	HayStockReturn      bool
	DescontarLlamadoCon string
}

func (s *StockFake) HayStock(libro string) bool {
	return s.HayStockReturn
}

func (s *StockFake) Descontar(libro string) {
	s.DescontarLlamadoCon = libro
}

// Mock de ProcesadorDePago
type PagoFake struct {
	CobrarReturn bool
}

func (p *PagoFake) Cobrar(monto float64) bool {
	return p.CobrarReturn
}

func TestComprar(t *testing.T) {
	casos := []struct {
		nombreCaso            string
		libro			  string
		hayStock          bool
		pagoOK            bool
		resultadoEsperado string
		esperaDescuento   bool
	}{
		{"Sin stock", "El Principito", false, false, "No hay stock", false},
		{"Pago rechazado", "El Principito", true, false, "Pago rechazado", false},
		{"Compra exitosa", "El Principito", true, true, "Compra exitosa", true},
	}

	for _, c := range casos {
		t.Run(c.nombreCaso, func(t *testing.T) {
			// Preparamos mocks
			stock := &StockFake{HayStockReturn: c.hayStock}
			pago := &PagoFake{CobrarReturn: c.pagoOK}

			// Capturar salida
			original := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Ejecutar Comprar(...)
			Comprar(c.libro, 100.0, stock, pago)

			// Leer la salida
			w.Close()
			os.Stdout = original
			out, _ := io.ReadAll(r)
			salida := string(out)

			// Verificar mensaje esperado
			if !strings.Contains(salida, c.resultadoEsperado) {
				t.Errorf("Esperaba \"%q\", pero se imprimio: %q", c.resultadoEsperado, salida)
			}

			// Verificar si Descontar fue llamado (cuando corresponde)
			if c.esperaDescuento && stock.DescontarLlamadoCon != c.libro {
				t.Errorf("Descontar fue llamado con %s, quiero %s", stock.DescontarLlamadoCon, c.libro)
			}
			if !c.esperaDescuento && stock.DescontarLlamadoCon != "" {
				t.Errorf("No esperaba que se llame a Descontar(), pero se llamo con: %q", stock.DescontarLlamadoCon)
			}
		})
	}
}
