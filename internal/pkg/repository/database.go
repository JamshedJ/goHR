package repository

import (
	"context"
	"fmt"

	"github.com/JamshedJ/goHR/internal/configs"
	"github.com/JamshedJ/goHR/internal/log"
	"github.com/jackc/pgx/v5"
)

type DB struct {
	conn *pgx.Conn
}

func New(ctx context.Context) *DB {
	cfg := configs.Database
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Address,
		cfg.Port,
		cfg.Name,
	)
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
