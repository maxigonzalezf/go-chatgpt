package main

import (
    "errors"
    "fmt"
)

// 1. Define el struct CuentaBancaria
type CuentaBancaria struct {
    Dueño  string
    Saldo  float64
}

// 2. Método Depositar: suma al saldo, recibe amount float64, no retorna error
func (c *CuentaBancaria) Depositar(amount float64) {
	c.Saldo += amount
}

// 3. Método Retirar: resta del saldo, recibe amount float64, retorna error si amount > Saldo
func (c *CuentaBancaria) Retirar(amount float64) error {
	if amount > c.getSaldo() {
		return errors.New("saldo insuficiente")
	}
	c.Saldo -= amount
	return nil
}

// 4. Método Saldo: retorna el saldo actual
func (c CuentaBancaria) getSaldo() float64 {
	return c.Saldo
}

// 5. Metodo nombre: retorna el nombre del dueño de la cuenta
func (c CuentaBancaria) getNombre() string {
	return c.Dueño
}

func main() {
    // a) Crea una cuenta con dueño "Tú" y saldo inicial 1000.50
    cuenta := CuentaBancaria{Dueño: "Tú", Saldo: 1000.50}
	fmt.Println(cuenta.getSaldo())

    // b) Haz un depósito de 250.75 e imprime el saldo
	cuenta.Depositar(250.75)
	fmt.Println(cuenta.getSaldo())

    // c) Intenta retirar 500 e imprime el saldo tras la operación
    if err := cuenta.Retirar(500); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Operacion exitosa!")
		fmt.Printf("Su saldo actual es de: %.2f\n", cuenta.getSaldo())
	}

    // d) Intenta retirar 2000 y maneja el error (mensaje informativo)
    if err := cuenta.Retirar(2000); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Operacion exitosa!")
		fmt.Printf("Su saldo actual es de: %.2f", cuenta.getSaldo())
	}

    // e) Imprime finalmente el saldo con cuenta.Saldo()
    fmt.Printf("Saldo final de %s: %.2f\n", cuenta.getNombre(), cuenta.getSaldo())
}
