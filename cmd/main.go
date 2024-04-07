package main

import (
	"context"
	"db-arch-lab2/internal/faker"
	"db-arch-lab2/internal/migrations"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

const DatabaseUrl = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
const ConnectionsCount = faker.WorkersCount

func main() {
	config, err := pgxpool.ParseConfig(DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	config.MaxConns = ConnectionsCount

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Unable to acquire a database connection: %v\n", err)
	}
	if err := migrations.Migrate("db/migration", conn.Conn()); err != nil {
		log.Fatal(err)
	}
	conn.Release()

	start := time.Now()
	faker.GenerateFakeData(pool)
	finish := time.Now()

	fmt.Println(finish.Sub(start))
}
