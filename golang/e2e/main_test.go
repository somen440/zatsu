package main

import (
	"context"
	"database/sql"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func TestStartServer(t *testing.T) {
	srv := buildServer()
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

	expect := "Hello, world!!"
	if string(b) != expect {
		t.Fatalf("want %q, but %q", expect, b)
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		t.Fatalf("HTTP server Shutdown: %v", err)
	}

	<-idleConnsClosed
}

func TestGetUserServer(t *testing.T) {
	var res *http.Response
	var err error

	srv := buildServer()
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		if err = srv.Serve(l); err != http.ErrServerClosed {
			t.Fatalf("HTTP server ListenAndServe: %v", err)
		}
		close(idleConnsClosed)
	}()

	retry := 5
	for {
		res, err = http.Get("http://" + l.Addr().String() + "/user")
		if err == nil {
			break
		}
		if retry == 0 {
			break
		}
		retry--
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()

	expect := "1:Olga:1978-01-06T15:27:33Z2:Scot:2002-07-19T04:24:42Z"
	if string(b) != expect {
		t.Fatalf("want %q, but %q", expect, b)
	}

	if err = srv.Shutdown(context.Background()); err != nil {
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

func dbConnect() (*sql.DB, error) {
	source := os.Getenv("USER") + ":" + os.Getenv("PASSWORD") + "@tcp(" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + ")/" + os.Getenv("DATABASE") + "?parseTime=true"
	return sql.Open("mysql", source)
}

type User struct {
	ID      int
	Name    string
	Created time.Time
}

func (u *User) ToString() string {
	return strconv.Itoa(u.ID) + ":" + u.Name + ":" + u.Created.Format(time.RFC3339)
}

func buildServer() *http.Server {
	r := mux.NewRouter()

	r.Methods("GET").
		Path("/").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, "Hello, world!!")
		})
	r.Methods("GET").
		Path("/user").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			db, err := dbConnect()
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
			defer db.Close()

			rows, err := db.Query("select * from users;")
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
			defer rows.Close()

			for rows.Next() {
				var user User
				err = rows.Scan(&user.ID, &user.Name, &user.Created)
				if err != nil {
					http.Error(w, err.Error(), 500)
				}
				io.WriteString(w, user.ToString())
			}
		})

	return &http.Server{
		Handler: r,
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
