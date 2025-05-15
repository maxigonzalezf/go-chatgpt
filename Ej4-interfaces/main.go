package main

import (
	"fmt"
)

// Interfaz Reportador: COMPLETAR
// Debe tener un método Reportar(mensaje string)
type Reportador interface {
	Reportar(string)
}


// Tipo EmailReportador
type EmailReportador struct {
	Direccion string
}

// Tipo ConsolaReportador
type ConsolaReportador struct{}

// FUNCIONES: Completá estos métodos

// 1. Método Reportar para EmailReportador
func (e EmailReportador) Reportar(mensaje string) {
	fmt.Printf("Enviando email a %s: %s\n", e.Direccion, mensaje)
}

// 2. Método Reportar para ConsolaReportador
func (c ConsolaReportador) Reportar(mensaje string) {
	fmt.Printf("Mostrando en consola: %s\n", mensaje)
}

// Función que recibe cualquier Reportador y envía un mensaje
func EnviarReporte(r Reportador, mensaje string) {
	r.Reportar(mensaje)
}

func main() {
	// COMPLETAR:
	// 1. Crear un EmailReportador con dirección "soporte@empresa.com"
	e := EmailReportador{"soporte@empresa.com"}
	// 2. Crear un ConsolaReportador
	c := ConsolaReportador{}
	// 3. Usar EnviarReporte con ambos para enviar un mensaje
	EnviarReporte(e, "Hola que tal")
	EnviarReporte(c, "Hola como estas")
}
