package main

import "github.com/fingernailz/assembler/internal/server"

func main() {
	server.CreateNewServer("8080").Run()
}
