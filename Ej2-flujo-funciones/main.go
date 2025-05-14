package main

import (
    "fmt"
)

func main() {
    nombres := []string{"Lucía", "Marcos", "Pedro", "Laura"}
    edades := []int{28, 17, 70, 0}

    // 1. Mostrar validaciones de edad
    for i := 0; i < len(nombres); i++ {
        nombre := nombres[i]
        edad := edades[i]

        if !esEdadValida(edad) {
            fmt.Printf("Edad inválida para %s\n", nombre)
            continue
        }

        estado := clasificarEdad(edad)
        fmt.Printf("%s tiene %d años (%s)\n", nombre, edad, estado)
    }

    // 2. Calcular el promedio de edad (de los válidos)
    promedio := calcularPromedio(edades)
    fmt.Printf("Promedio de edad (válidos): %.2f\n", promedio)
}

// esEdadValida devuelve true si edad está entre 1 y 120
func esEdadValida(edad int) bool {
	if edad > 0 && edad < 121 {
		return true
	}

	return false
}

// clasificarEdad devuelve un string según la edad
// <18: "Menor", 18–65: "Adulto", >65: "Jubilado"
func clasificarEdad(edad int) string {
	estado := "Menor"
	switch {
		case edad > 18 && edad <= 65:
			estado = "Adulto"
		case edad > 65:
			estado = "Jubilado"
	}
	return estado
}

// calcularPromedio calcula promedio de edades válidas
func calcularPromedio(edades []int) float64 {
	total := 0.0
	invalidos := 0
	for _, edad := range edades {
		if !esEdadValida(edad) {
			invalidos++
			continue
		}
		total += float64(edad)
	}
	promedio := total / float64(len(edades) - invalidos)

	return float64(promedio)
}
