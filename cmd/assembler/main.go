package main

import (
	"log"

	"github.com/fingernailz/assembler/internal/assembler"
)

func main() {
	asm := assembler.CreateAssemblerFile("what.asm")
	log.Println(asm.Assemble())
}
