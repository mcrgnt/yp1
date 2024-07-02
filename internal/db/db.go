package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func getDB(dsn string) (conn *pgx.Conn, err error) {
	if conn, err = pgx.Connect(context.Background(), dsn); err != nil {
		err = fmt.Errorf("db connect failed: %w", err)
		return
	}
	return nil, nil
}

func freeDB(conn *pgx.Conn) (err error) {
	if err = conn.Close(context.Background()); err != nil {
		err = fmt.Errorf("db close failed: %w", err)
	}
	return
}

func Ping(dsn string) (err error) {
	var (
		conn *pgx.Conn
	)
	if conn, err = getDB(dsn); err != nil {
		err = fmt.Errorf("get db failed: %w", err)
		return
	}
	defer func() {
		if e := freeDB(conn); e != nil {
			err = errors.Join(err, fmt.Errorf("db close failed: %w", e))
		}
	}()
	if err = conn.Ping(context.Background()); err != nil {
		err = fmt.Errorf("db ping failed: %w", err)
		return
	}
	return nil
}
