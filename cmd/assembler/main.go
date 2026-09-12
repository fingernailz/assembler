package main

import (
	"log"

	"github.com/fingernailz/assembler/internal/assembler"
)

func main() {
	asm := assembler.CreateAssembler("what.asm")
	log.Println(asm.Assemble())
}
