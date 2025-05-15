package main

import (
	"fmt"
)

// Interfaz MedioDePago
// COMPLETAR: debe tener un método Pagar(monto float64) bool
type MedioDePago interface {
	Pagar(monto float64) bool
	ObtenerTitular() string
}

// Tipo TarjetaCredito
type TarjetaCredito struct {
	NombreTitular string
	Numero        string
	Fondos        float64
}

// Tipo CuentaBancaria
type CuentaBancaria struct {
	NombreTitular string
	Alias         string
	Saldo         float64
}

// IMPLEMENTACIÓN de los métodos para ambos tipos:
// COMPLETAR...
func (tc *TarjetaCredito) Pagar(monto float64) bool {
	if tc.Fondos >= monto {
		tc.Fondos -= monto
		fmt.Printf("Pagando con la tarjeta %s...\n", tc.Numero)
		return true
	}
	return false
}

func (cb *CuentaBancaria) Pagar(monto float64) bool {
	if cb.Saldo >= monto {
		cb.Saldo -= monto
		fmt.Printf("Pagando con la cuenta: %s...\n", cb.Alias)
		return true
	}
	return false
}

func (tc TarjetaCredito) ObtenerTitular() string {
	return fmt.Sprintf("con la tarjeta de %s", tc.NombreTitular)
}

func (cb CuentaBancaria) ObtenerTitular() string {
	return fmt.Sprintf("con la cuenta de %s", cb.NombreTitular)
}

// Función que recibe cualquier MedioDePago y procesa un pago
func ProcesarPago(m MedioDePago, monto float64) {
	fmt.Printf("Intentando cobrar $%.2f %s\n", monto, m.ObtenerTitular())
	if m.Pagar(monto) {
		fmt.Println("Pago procesado con éxito ✅")
	} else {
		fmt.Println("Pago rechazado ❌")
	}
	fmt.Println()
}

func main() {
	// COMPLETAR:

	// 1. Crear una tarjeta con $500
	tc := &TarjetaCredito{"Norma", "xxxx", 500}
	// 2. Crear una cuenta bancaria con $300
	cb := &CuentaBancaria{"Maxi", "maxi.mp", 300}
	// 3. Intentar pagar $200 con ambos métodos
	ProcesarPago(tc, 200)
	ProcesarPago(cb, 200)
	// 4. Intentar pagar $400 con ambos métodos
	ProcesarPago(tc, 400)
	ProcesarPago(cb, 400)
}
