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

// internal/infrastructure/persistence/sql_repo.go (agregar esto)
func (r *PedidoRepoSQL) SaveLinea(lp domain.LineaPedido) error {
	_, err := r.db.Exec(`
		INSERT INTO lineas_pedido (id, pedido_id, producto, cantidad, moneda, precio)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		lp.ID, lp.PedidoID, lp.Producto, lp.Cantidad, lp.Precio.Moneda, lp.Precio.Cantidad)
	return err
}

func (r *PedidoRepoSQL) FindLineByPedidoID(pedidoID string) ([]domain.LineaPedido, error) {
	rows, err := r.db.QueryContext(context.Background(),
		`SELECT id, producto, cantidad, moneda, precio FROM lineas_pedido WHERE pedido_id = $1`, pedidoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lineas []domain.LineaPedido
	for rows.Next() {
		var (
			id       string
			producto string
			cantidad int
			moneda   string
			precio   float64
		)

		if err := rows.Scan(&id, &producto, &cantidad, &moneda, &precio); err != nil {
			return nil, err
		}

		lp := domain.LineaPedido{
			ID:       id,
			PedidoID: pedidoID,
			Producto: producto,
			Cantidad: cantidad,
			Precio: domain.Dinero{
				Moneda:   moneda,
				Cantidad: precio,
			},
		}

		lineas = append(lineas, lp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lineas, nil
}
