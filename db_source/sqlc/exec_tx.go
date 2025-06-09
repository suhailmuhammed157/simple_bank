package db_source

import (
	"context"
	"fmt"
)

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbError := tx.Rollback(ctx); rbError != nil {
			return fmt.Errorf("tx Error: %v, rb Error: %v", err, rbError)
		}
		return err
	}
	return tx.Commit(ctx)
}
