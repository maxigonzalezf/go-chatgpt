package main

import (
    "flag"
    "fmt"
    "os"
	"github.com/maxigonzalezf/todo-cli/todo"
)

func main() {
    // Definir comandos y flags:
    addCmd := flag.NewFlagSet("add", flag.ExitOnError)
    listCmd := flag.NewFlagSet("list", flag.ExitOnError)
    completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
    completeID := completeCmd.Int("id", 0, "ID de la tarea a completar")

    if len(os.Args) < 2 {
        fmt.Println("uso: todo-cli [add|list|complete]")
        os.Exit(1)
    }

    lista := todo.NuevaLista()

    switch os.Args[1] {
    case "add":
        addCmd.Parse(os.Args[2:])
        texto := addCmd.Arg(0)
        if texto == "" {
            fmt.Println("especificá el texto de la tarea")
            os.Exit(1)
        }
        lista.Add(texto)
        fmt.Println("tarea agregada")

    case "list":
        listCmd.Parse(os.Args[2:])
        for _, t := range lista.List() {
            status := " "
            if t.Complet {
                status = "x"
            }
            fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Texto)
        }

    case "complete":
        completeCmd.Parse(os.Args[2:])
        if *completeID == 0 {
            fmt.Println("especificá -id de la tarea")
            os.Exit(1)
        }
        if lista.Complete(*completeID) {
            fmt.Println("tarea completada")
        } else {
            fmt.Println("tarea no encontrada")
        }

    default:
        fmt.Println("comando no reconocido")
        os.Exit(1)
    }
}