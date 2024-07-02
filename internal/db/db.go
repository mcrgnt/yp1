package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Ping(dsn string) (err error) {
	var (
		conn *pgx.Conn
	)
	if conn, err = pgx.Connect(context.Background(), dsn); err != nil {
		err = fmt.Errorf("db connect failed: %w", err)
		return
	}
	defer func() {
		if e := conn.Close(context.Background()); e != nil {
			err = errors.Join(err, fmt.Errorf("db close failed: %w", e))
		}
	}()
	if err = conn.Ping(context.Background()); err != nil {
		err = fmt.Errorf("db ping failed: %w", err)
		return
	}
	return nil
}
