package assembler

import (
	"os"
	"strings"

	"github.com/fingernailz/assembler/internal/isa"
)

/*
	we allocate 64 bits for the label and 32 bits for normal instructiosn
*/

type Assembler struct {
	FileName    string //along with the extension
	SymbolTable map[string]int

	// stores the well instructions, instructionstable is a better name but lets go with something new
	StatementTable []string // I hope this is fine, dk what else data structrue to use for fix later ig????/
}

func (asm *Assembler) BuildSymbolicTable() {

	data, err := os.ReadFile(asm.FileName)

	if err != nil {
		panic(err)
	}

	stringData := string(data)

	splitStringData := strings.Split(stringData, "\n")

	asm.CreateFINDNAME(splitStringData)

}

// idk what I have to name this function as!!!!
func (asm *Assembler) CreateFINDNAME(programArray []string) {
	var locationCounter int = 0

	for _, line := range programArray {

		//strip the line of spaces first
		line = strings.TrimSpace(line)

		//idk why but doing it anyways

		if line == "" {
			continue
		}

		// for having a valid label shoudl I need to follow literal naming convention or something else?
		// as of now I will just see if the label is not named after any instruction, ends with : and is not repeated again via the symbol table

		//change the hard coded value to a const
		if !strings.HasSuffix(line, ":") {
			asm.StatementTable = append(asm.StatementTable, line)
			locationCounter++
			continue
		}

		if _, ok := asm.SymbolTable[line]; !ok {
			panic("same label has been ridden twice")
		}

		// something like "asd DFdf:" should throw error and "adf :" should not idk how to do this, lets rawdog this
		if xs := strings.Split(line, " "); len(xs) != 1 && xs[1] != ":" {
			panic("errors name that I couldnt explain")
		}

		//checks if the word is an instruction from the isa package
		label := strings.TrimRight(line, ":")

		if _, ok := isa.Instructions[label]; ok {
			panic("The label is an instruction")
		}

		asm.SymbolTable[label] = locationCounter

		locationCounter++

	}

}
