package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

type DB struct {
	conn *pgx.Conn
}

func New(ctx context.Context, dsn string) *DB {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("Enable connect to database: ", err)
	}
	return &DB{conn: conn}
}

func (d *DB) Close(ctx context.Context) {
	if d.conn == nil {
		return
	}
	if err := d.conn.Close(ctx); err != nil {
		log.Fatal("Error closing database: ", err)
	}
}
