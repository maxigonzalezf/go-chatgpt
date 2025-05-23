package persistence

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/maxigonzalezf/go-chatgpt/Ej7-microservicio-pedidos/internal/domain"
)

type PedidoRepoSQL struct {
    db *sql.DB
}

// Nuevo adaptador SQL, recibe un *sql.DB configurado
func NewPedidoRepoSQL(db *sql.DB) *PedidoRepoSQL {
    return &PedidoRepoSQL{db: db}
}

func (r *PedidoRepoSQL) Save(p domain.Pedido) error {
    // La tabla pedidos guarda cada pedido como una fila
	// Hace INSERT ... ON CONFLICT para upsert
	query := `INSERT INTO pedidos (id, moneda, cantidad) VALUES ($1, $2, $3)
              ON CONFLICT (id) DO UPDATE SET moneda = EXCLUDED.moneda, cantidad = EXCLUDED.cantidad`
    // ExecContext con context.Background()
	_, err := r.db.ExecContext(context.Background(), query, p.ID, p.Total.Moneda, p.Total.Cantidad)
    return err
}

func (r *PedidoRepoSQL) FindByID(id string) (domain.Pedido, error) {
    var moneda string
    var cantidad float64

    query := `SELECT moneda, cantidad FROM pedidos WHERE id = $1`
    // QueryRowContext con context.Background()
	row := r.db.QueryRowContext(context.Background(), query, id)
    if err := row.Scan(&moneda, &cantidad); err != nil {
        if err == sql.ErrNoRows {
            return domain.Pedido{}, fmt.Errorf("pedido no encontrado")
        }
        return domain.Pedido{}, err
    }

    return domain.Pedido{
        ID: id,
        Total: domain.Dinero{
            Moneda:   moneda,
            Cantidad: cantidad,
        },
    }, nil
}
