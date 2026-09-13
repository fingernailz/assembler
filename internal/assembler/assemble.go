package assembler

import (
	"os"
	"strings"

	"github.com/fingernailz/assembler/internal/formats"
	"github.com/fingernailz/assembler/internal/isa"
)

const (
	labelSuffix string = ":"
	newLine     string = "\n"
	space       string = " "
	emptyString string = ""
)

/*
	we allocate 64 bits for the label and 32 bits for normal instructiosn
*/

type Assembler struct {
	//privating this
	fileName    string //along with the extension
	SymbolTable map[string]int

	// stores the well instructions, instructionstable is a better name but lets go with something new
	// lol its just a array not a table
	StatementTable   []string // I hope this is fine, dk what else data structrue to use for fix later ig????/
	InstructionArray []string // all the lines
}

func (asm *Assembler) buildSymbolicTable() {

	data, err := os.ReadFile(asm.fileName)

	if err != nil {
		panic(err)
	}

	stringData := string(data)

	splitStringData := strings.Split(stringData, newLine)

	asm.InstructionArray = splitStringData

	asm.CreateFINDNAME()

}

// idk what I have to name this function as!!!!
func (asm *Assembler) CreateFINDNAME() {
	var locationCounter int = 0

	programArray := asm.InstructionArray

	if len(programArray) == 0 {
		panic("error with program error")
	}

	for _, line := range programArray {

		//strip the line of spaces first
		line = strings.TrimSpace(line)

		//idk why but doing it anyways

		if line == emptyString {
			continue
		}

		// for having a valid label shoudl I need to follow literal naming convention or something else?
		// as of now I will just see if the label is not named after any instruction, ends with : and is not repeated again via the symbol table

		//change the hard coded value to a const
		if !strings.HasSuffix(line, labelSuffix) {
			asm.StatementTable = append(asm.StatementTable, line)
			locationCounter++
			continue
		}

		// redundency label
		if _, ok := asm.SymbolTable[line]; !ok {
			panic("same label has been ridden twice")
		}

		// something like "asd DFdf:" should throw error and "adf :" should not idk how to do this, lets rawdog this
		if xs := strings.Split(line, space); len(xs) != 1 && xs[1] != labelSuffix {
			panic("errors name that I couldnt explain")
		}

		//checks if the word is an instruction from the isa package
		label := strings.TrimRight(line, labelSuffix)

		if _, ok := isa.Instructions[label]; ok {
			panic("The label is an instruction")
		}

		// the label will point towards the next insturction but we don't update the locationcounter itself if we find a valid label
		asm.SymbolTable[label] = locationCounter + 1

	}

}

//labels are not encoded in machine code, we go to the next insturction present in the label
// for jumps. fix this in the function

// again lol find perfect or smoewhat perfect name
func (asm *Assembler) pass2() string {
	finalIns, err := formats.CreateFormats(asm.StatementTable, &asm.SymbolTable)

	if err != nil {
		panic(err)
	}

	var outputBinary strings.Builder
	var temp string

	for _, y := range finalIns.InstructionSet {
		temp = y.ConvertToBinary()
		outputBinary.WriteString(temp) // for loop wrt += is inefficient itseems idk chekc later
	}

	return outputBinary.String()
}

func CreateAssemblerFile(filename string) *Assembler {

	if filename == emptyString {
		panic("error: empty string and data")
	}

	return &Assembler{
		fileName:         filename,
		InstructionArray: nil,
	}
}

func CreateAssemblerServer(data []string) *Assembler {

	if len(data) == 0 {
		return nil
	}

	return &Assembler{
		InstructionArray: data,
		fileName:         emptyString,
	}
}

// this returns the final string of binary data
func (asm *Assembler) Assemble() string {
	if asm.fileName == emptyString {
		asm.buildSymbolicTable()
	} else {
		asm.CreateFINDNAME()
	}

	return asm.pass2()
}
