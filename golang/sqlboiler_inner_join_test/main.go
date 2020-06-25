package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("-- start")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	db, err := sql.Open("mysql", "test_user:test_pass@tcp(localhost:3232)/test_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		return err
	}
	defer db.Close()

	entities, err := SelectAuthorPosts(db, ctx, 1, 2, 3)
	if err != nil {
		return err
	}

	for _, entity := range entities {
		fmt.Println(entity)
	}
	return nil
}
