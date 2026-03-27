package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

// func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello from the server"))
// 	switch r.Method {
// 	case http.MethodGet:
// 		switch r.URL.Path {
// 		case "/":
// 			w.Write([]byte("index"))
// 			return
// 		case "/lmao":
// 			w.Write([]byte("lmao"))
// 			return
// 		}
// 	default:
// 		w.Write([]byte("404"))
// 	}
// }

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}

func main() {
	s := &api{addr: ":8080"}

	mux := http.NewServeMux() // Basically for creating router

	srv := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", s.getUsersHandler) 

	log.Fatal(srv.ListenAndServe())
}
