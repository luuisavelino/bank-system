package repository

import (
	"context"

	_ "github.com/lib/pq"
)

// CheckIfClientExists check if client exists
func (sr bankRepository) CheckIfClientExists(ctx context.Context, clientId int64) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM clients WHERE id = $1)`

	err := sr.db.QueryRow(ctx, query, clientId).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
