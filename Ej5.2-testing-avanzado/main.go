package main

import "fmt"

// Interfaces
type StockChecker interface {
    HayStock(libro string) bool
    Descontar(libro string)
}

type ProcesadorDePago interface {
    Cobrar(monto float64) bool
}

// Función principal: procesar una compra
func Comprar(libro string, precio float64, s StockChecker, p ProcesadorDePago) {
    if !s.HayStock(libro) {
        fmt.Println("No hay stock 😓")
        return
    }

    if !p.Cobrar(precio) {
        fmt.Println("Pago rechazado 💸")
        return
    }

    s.Descontar(libro)
    fmt.Println("Compra exitosa 🎉")
}

func main(){}
