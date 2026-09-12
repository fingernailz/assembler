package formats

import (
	"strings"

	"github.com/fingernailz/assembler/internal/isa"
	"github.com/fingernailz/assembler/internal/models"
)

const (
	emptyString string = ""
	space       string = " "
)

// rename this type later
type NewTyep struct {
	InstructionSet []*models.Instruction
	SymbolTablePtr *map[string]int
}

func CreateFormats(instructions []string, stptr *map[string]int) *NewTyep {

	for _, instr := range instructions {

		// doing this again in second pass
		instr = strings.TrimSpace(instr)

		if instr == emptyString {
			continue
		}

		instrSlice := strings.Split(instr, " ")

		mappedValue, ok := isa.Instructions[strings.ToUpper(instrSlice[0])]

		if !ok {
			//errors
			// return nil, new errors("error")
		}

		switch mappedValue["format"] {
		case isa.R_FORMAT:
		case isa.I_FORMAT:
		case isa.D_FORMAT:
		case isa.B_FORMAT:
		case isa.CB_FORMAT:
		case isa.IW_FORMAT:
		}
	}

	return nil
}
