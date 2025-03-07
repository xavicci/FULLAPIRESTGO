package main

import (
	"fmt"
	"net/http"
	"time"
)

type HelloHandler struct{}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("heeello Ecuador!"))
}

func main() {

	person := http.NewServeMux()
	person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("greetings!\n"))
	})

	dog := http.NewServeMux()
	dog.HandleFunc("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		w.Write([]byte(fmt.Sprintf("Hello good puppy, %s!\n", name)))
	})

	mux := http.NewServeMux()

	mux.HandleFunc("GET /inicio/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		w.Write([]byte(fmt.Sprintf("Hello, %s!\n", name)))

	})

	mux.Handle("/iniciosinhandle", HelloHandler{})

	mux.Handle("/person/", http.StripPrefix("/person", person))
	mux.Handle("/dog/", http.StripPrefix("/dog", dog))

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
