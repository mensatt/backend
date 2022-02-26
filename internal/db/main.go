//go:generate go run github.com/kyleconroy/sqlc/cmd/sqlc generate

package db

// // Repository defines all functions to execute db queries and transactions
// type Repository interface {
// 	Querier
// }

// type repoSvc struct {
// 	*Queries
// 	db *pgxpool.Pool
// }

// // NewRepository returns an implementation of the Repository interface.
// func NewRepository(db *pgxpool.Pool) Repository {
// 	return &repoSvc{
// 		Queries: New(db),
// 		db:      db,
// 	}
// }

// func (r *repoSvc) withTx(ctx context.Context, txFn func(*Queries) error) error {
// 	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
// 	if err != nil {
// 		return err
// 	}
// 	q := New(tx)
// 	err = txFn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(ctx); rbErr != nil {
// 			return fmt.Errorf("tx failed: %v, unable to rollback: %v", err, rbErr)
// 		}
// 		return err
// 	}
// 	return tx.Commit(ctx)
// }
