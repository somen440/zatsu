package main

import (
	"context"
	"database/sql"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
)

func buildServer(body string) *http.Server {
	r := mux.NewRouter()

	r.Methods("GET").
		PathPrefix("/").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, body)
		})

	return &http.Server{
		Handler: r,
	}
}

func TestStartServer(t *testing.T) {
	want := "Hello, world!\n"

	srv := buildServer(want)
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		if err := srv.Serve(l); err != http.ErrServerClosed {
			t.Fatalf("HTTP server ListenAndServe: %v", err)
		}
		close(idleConnsClosed)
	}()

	res, err := http.Get("http://" + l.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()

	if string(b) != want {
		t.Fatalf("want %q, but %q", want, b)
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		t.Fatalf("HTTP server Shutdown: %v", err)
	}

	<-idleConnsClosed
}

func TestCreateItemWithDBMock(t *testing.T) {
	name := "hoge"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("intert into items").
		WithArgs(name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err := createItem(db, name); err != nil {
		t.Errorf(err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf(err.Error())
	}
}

func createItem(db *sql.DB, name string) error {
	var err error

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	query := "intert into items (name) values (?)"
	if _, err = tx.Exec(query, name); err != nil {
		return err
	}

	return nil
}
