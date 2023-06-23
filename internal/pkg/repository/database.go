package repository

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/jackc/pgx/v5"
)

type DB struct {
	conn *pgx.Conn
}

func New(ctx context.Context, dsn string) *DB {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Error.Fatal("error connecting to database: ", err)
	}

	log.Debug.Println("database connected")
	return &DB{conn: conn}
}

func (d *DB) Close(ctx context.Context) {
	if d.conn == nil {
		log.Warning.Println("db.Close(): database connection already nil")
		return
	}
	if err := d.conn.Close(ctx); err != nil {
		log.Error.Fatal("error closing database: ", err)
	}
	log.Debug.Println("database connection closed")
}
