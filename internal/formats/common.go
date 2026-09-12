package formats

import (
	"strings"

	"github.com/fingernailz/assembler/internal/isa"
	"github.com/fingernailz/assembler/internal/models"
)

const (
	emptyString string = ""
	space       string = " "
	comma       string = ","
)

// rename this type later
type NewTyep struct {
	// again no need for *models.Instruction
	InstructionSet []models.Instruction
	SymbolTablePtr *map[string]int
}

// correct name, nope change
func CreateFormats(instructions []string, stptr *map[string]int) (*NewTyep, error) {

	formatType := NewTyep{}
	formatType.SymbolTablePtr = stptr // this is for passing to the parser of certain formats

	var (
		ptr models.Instruction
		err error
	)

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
			// a normal pointer satisfies models.Instruction
			ptr, err = RFCreate(instr)
		case isa.I_FORMAT:
		case isa.D_FORMAT:
		case isa.B_FORMAT:
		case isa.CB_FORMAT:
		case isa.IW_FORMAT:
		}

		if err != nil {
			return nil, err
		}

		formatType.InstructionSet = append(formatType.InstructionSet, ptr)

	}

	return &formatType, err
}
