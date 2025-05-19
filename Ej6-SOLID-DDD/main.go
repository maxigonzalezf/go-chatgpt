package main

import (
	"fmt"

	"github.com/maxigonzalezf/go-chatgpt/Ej6-SOLID-DDD/dinero"
	"github.com/maxigonzalezf/go-chatgpt/Ej6-SOLID-DDD/pedido"
)

// Value Object Dinero, tiene valor pero no identidad, incluye logica de suma y validacion
// Entidad Pedido con identidad (ID) y comportamiento (AgregarLinea)
// Interfaz de repositorio PedidoRepository (port), abstrae el guardado y lectura de pedidos
// Implementacion concreta en memoria del repo (adapter)
// Funcion de servicio de dominio AgregarLinea() con una regla de negocio (no permite lineas en distintas monedas)

func main() {

	repo := pedido.NuevoRepositorioMemoria() // inyeccion de dependencias
	p := pedido.Pedido{ID: "P-001"}

	err := pedido.AgregarLinea(&p, dinero.Dinero{Moneda: "USD", Cantidad: 200})
	if err != nil {
		fmt.Println("Error al agregar linea 1:", err)
	}

	err = pedido.AgregarLinea(&p, dinero.Dinero{Moneda: "USD", Cantidad: 400})
	if err != nil {
		fmt.Println("Error al agregar linea 2:", err)
	}

	err = pedido.AgregarLinea(&p, dinero.Dinero{Moneda: "Real", Cantidad: 100})
	if err != nil {
		fmt.Println("Error esperado (moneda distinta):", err)
	}

	_ = repo.Save(p)

	guardado, err := repo.FindByID("P-001")
	if err != nil {
		fmt.Println("Error al buscar pedido:", err)
	} else {
		fmt.Printf("Pedido guardado: ID = %s | Total = %.2f %s\n",
			guardado.ID, guardado.Total.Cantidad, guardado.Total.Moneda)
	}
}