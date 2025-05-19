package main

import "fmt"

// MedioDePago define el contrato para procesar pagos.
type MedioDePago interface {
	Pagar(monto float64) bool
}

// ProcesarPago invoca Pagar y escribe en consola el resultado.
func ProcesarPago(m MedioDePago, monto float64) {
	fmt.Printf("Intentando cobrar $%.2f...\n", monto)
	if m.Pagar(monto) {
		fmt.Println("Pago procesado con éxito ✅")
	} else {
		fmt.Println("Pago rechazado ❌")
	}
}

func main() {

}
