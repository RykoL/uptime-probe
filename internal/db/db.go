package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func CreateDBPool(connectionString string) (*pgxpool.Pool, error) {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	return dbpool, nil
}
