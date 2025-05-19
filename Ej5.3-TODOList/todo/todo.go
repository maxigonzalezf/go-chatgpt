package todo

// Tarea representa una tarea pendiente.
type Tarea struct {
    ID      int
    Texto   string
    Complet bool
}

// Lista mantiene todas las tareas y el siguiente ID libre.
type Lista struct {
    tareas   []Tarea
    siguiente int
}

// NuevaLista inicializa una lista vacía.
func NuevaLista() *Lista {
    return &Lista{
        tareas:    []Tarea{},
        siguiente: 1,
    }
}

// Add agrega una nueva tarea con texto dado.
func (l *Lista) Add(texto string) {
    // TODO: crear una Tarea con l.siguiente, texto y Complet=false,
    //       luego añadirla a l.tareas y aumentar l.siguiente.
	tarea := Tarea{ID: l.siguiente, Texto: texto, Complet: false}
	l.tareas = append(l.tareas, tarea)
	l.siguiente += 1
}

// List devuelve el slice completo de tareas.
func (l *Lista) List() []Tarea {
    // TODO: retornar l.tareas
    return l.tareas
}

// Complete marca como completada la tarea con el ID dado.
// Si no existe, devolver false; si la marca, devolver true.
func (l *Lista) Complete(id int) bool {
    // TODO: buscar en l.tareas la tarea con ID == id,
    //       si la encontrás marca Complet=true y retorna true,
    //       de lo contrario retorna false.
	for i := range l.tareas {
		if l.tareas[i].ID == id {
			l.tareas[i].Complet = true
			return true
		}
	}
    return false
}