// OBJETIVO
// Practicar:
// 	- Declaracion de variables y constantes
//	- Uso de arrays, slices y maps
//	- Tipos basicos
//	- Sintaxis general de Go

package main

import (
	"fmt"
)

func main() {
	var names = []string{"Mara", "Maxi"}
	addPerson("Norma", &names)
	fmt.Println("NOMBRES")
	showList(names)

	ages := []int{28, 29}
	addAge(64, &ages)
	fmt.Println("EDADES")
	showList(ages)
	avg := ageAvg(ages)
	fmt.Printf("El promedio es: %.2f", avg)
	fmt.Println()

	status := []bool{true, false}
	addStatus(true, &status)
	fmt.Println("STATUS")
	showList(status)

	personsMap := map[string]int{
		"Mara":  28,
		"Maxi":  29,
		"Norma": 64,
	}
	fmt.Println("MAPA")
	printMap(personsMap)

	age := searchAgeofPerson("Norma", personsMap)

	if age != 0 {
		fmt.Printf("La edad es: %v", age)
	} else {
		fmt.Printf("La persona solicitada no se encuentra en la lista")
	}

	addPersonToMap("Mario", 64, &personsMap)
	fmt.Println()
	printMap(personsMap)
}

// CREAR UN PROGRAMA QUE DEFINA INFO DE PERSONAS
func addPerson(name string, personsList *[]string) {
	*personsList = append(*personsList, name)
}

func addAge(age int, agesList *[]int) {
	*agesList = append(*agesList, age)
}

func addStatus(status bool, statusList *[]bool) {
	*statusList = append(*statusList, status)
}

// MOSTRAR LOS DATOS
func showList[T any](anyList []T) {
	for i, v := range anyList {
		fmt.Printf("%v) %v", (i + 1), v)
		fmt.Println()
	}

}

// CALCULAR PROMEDIO DE EDADES
func ageAvg(agesList []int) float64 {
	total := 0
	for _, v := range agesList {
		total += v
	}

	return (float64)(total) / (float64)(len(agesList))
}

// BUSCAR LA EDAD POR NOMBRE CON MAPS
func searchAgeofPerson(name string, personsMap map[string]int) int {
	age, exists := personsMap[name]
	if !exists {
		return 0
	}

	return age
}

// BONUS
func printMap(personsMap map[string]int) {
	for k, v := range personsMap {
		fmt.Println(k, ":", v)
	}
}
func addPersonToMap(name string, age int, personsMap *map[string]int) {
	(*personsMap)[name] = age
}
