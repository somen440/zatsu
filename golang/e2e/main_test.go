package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"testing"

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
