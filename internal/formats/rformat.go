package formats

import (
	"strings"

	"github.com/fingernailz/assembler/internal/isa"
	"github.com/fingernailz/assembler/internal/registers"
)

type RFromat struct {
	Instruction         string
	FirstRegister       string
	SecondRegister      string
	DestinationRegister string
}

// eg. ADD X0, X1, X2
// split this by space and remove the comma
func RFCreate(instruction string) (*RFromat, error) {

	instruction = strings.TrimSpace(instruction)
	insSlice := strings.Split(instruction, space)

	// what the fuck, why is this shit soo hard

	if len(insSlice) != 4 {
		return nil, nil // TODO: change to appropiate error
	}

	out, ok := isa.Instructions[strings.ToUpper(insSlice[0])]

	if !ok {
		return nil, nil // TODO: change to appropiate error
	}

	form := RFromat{}

	opcode := out["op-code"]

	form.Instruction = opcode

	for x, y := range insSlice[1:] {

		//lmafoo alr its stupid as hell
		insSlice[x] = strings.TrimSpace(strings.Trim(strings.TrimSpace(y), comma))
		value, ok := registers.RegistersBin[insSlice[x]]

		if !ok {
			return nil, nil // error agin
		}

		func(x int) {
			switch x {
			case 1:
				form.DestinationRegister = value
			case 2:
				form.FirstRegister = value
			default:
				form.SecondRegister = value
			}
		}(x)

	}

	return &form, nil
}

func (rf *RFromat) ConvertToBinary() string {
	binary := ""

	return binary
}
