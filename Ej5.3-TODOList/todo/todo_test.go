package todo

import "testing"

func TestAddYList(t *testing.T) {
    l := NuevaLista()
    l.Add("Comprar pan")
    l.Add("Llamar a mamá")

    tareas := l.List()
    if len(tareas) != 2 {
        t.Errorf("esperaba 2 tareas, tengo %d", len(tareas))
    }
    if tareas[0].Texto != "Comprar pan" {
        t.Errorf("tarea 1 texto = %q, quiero %q", tareas[0].Texto, "Comprar pan")
    }
    if tareas[1].ID != 2 {
        t.Errorf("tarea 2 ID = %d, quiero %d", tareas[1].ID, 2)
    }
}

func TestComplete(t *testing.T) {
    l := NuevaLista()
    l.Add("Estudiar Go")
    ok := l.Complete(1)
    if !ok {
        t.Errorf("esperaba que Complete(1) devolviera true")
    }
    if !l.List()[0].Complet {
        t.Errorf("esperaba que la tarea 1 quedara marcada como Complet")
    }

    notFound := l.Complete(999)
    if notFound {
        t.Errorf("esperaba que Complete(999) devolviera false")
    }
}