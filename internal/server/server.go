package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fingernailz/assembler/internal/assembler"
)

type Server struct {
	*http.ServeMux
	Port string
}

func CreateNewServer(port string) *Server {
	_, err := strconv.Atoi(port)

	if err != nil {
		panic("error")
	}

	servermux := http.NewServeMux()

	servermux.HandleFunc(
		"GET /{$}",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hii"))
			w.WriteHeader(http.StatusAccepted)
		},
	)

	servermux.HandleFunc(
		"GET /assemble",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("new request from somewhere")
			// testing as of now
			asm := assembler.CreateAssemblerServer([]string{"ADD X1, X2, X3"})
			asdf := asm.Assemble()
			w.Write([]byte(asdf))
		},
	)

	return &Server{
		Port:     port,
		ServeMux: servermux,
	}
}

func (S *Server) Run() {
	log.Println("Server running at port", S.Port)
	log.Fatal(http.ListenAndServe(":"+S.Port, S.ServeMux))
}
