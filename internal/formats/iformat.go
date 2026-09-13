package formats

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/fingernailz/assembler/internal/isa"
	"github.com/fingernailz/assembler/internal/registers"
)

const (
	LOWER_BOUND             int    = 0
	UPPER_BOUND             int    = 4095
	BASE_2                  int    = 2
	ZERO_APPEND             string = "0"
	IMMEDIATE_STRING_LENGTH int    = 12
)

type IFormat struct {
	Instruction         string
	DestinationRegister string
	FirstRegister       string
	ImmediateValue      string // max value is 2 ^ 12, 0 append front. 4,095
}

func (iform *IFormat) ConvertToBinary() string {
	binaryOutput := iform.Instruction + iform.ImmediateValue + iform.DestinationRegister + iform.FirstRegister
	return binaryOutput
}

// eg. ADDI X9, X9, #1
// I need to split by space, remove ',', remove # and parse the string to int
// what if ADDDI x8,x8,#1
// idk I will do that later not now
func IFCreate(instruction string) (*IFormat, error) {

	instruction = strings.TrimSpace(instruction)
	insSlice := strings.Split(instruction, space)

	if len(insSlice) != 4 {
		return nil, nil
	}

	out, ok := isa.Instructions[strings.ToUpper(insSlice[0])]

	if !ok {
		return nil, nil
	}

	form := IFormat{}

	opcode := out["op-code"]

	form.Instruction = opcode

	dr := strings.TrimSpace(
		strings.Trim(
			strings.TrimSpace(insSlice[1]), comma,
		),
	)

	value, ok := registers.RegistersBin[strings.ToUpper(dr)]

	if !ok {
		panic("error")
	}

	form.DestinationRegister = value

	//comma should be included

	if !(strings.HasSuffix(insSlice[1], comma) && strings.HasSuffix(insSlice[2], comma)) {
		panic("ohh error")
	}

	fr := strings.TrimSpace(
		strings.Trim(
			strings.TrimSpace(insSlice[2]), comma,
		),
	)

	value, ok = registers.RegistersBin[strings.ToUpper(fr)]

	if !ok {
		panic("error 1")
	}

	form.FirstRegister = value

	insSlice[3] = strings.TrimSpace(insSlice[3])

	if !strings.HasPrefix(insSlice[3], HASHTAG) {
		panic("error 2")
	}

	imme := strings.TrimLeft(insSlice[3], HASHTAG)

	intImmediate, err := strconv.Atoi(imme)

	if err != nil {
		log.Println("error")
		return nil, nil
	}

	if intImmediate < LOWER_BOUND || intImmediate > UPPER_BOUND {
		panic("error 3")
	}

	bitString := strconv.FormatInt(int64(intImmediate), BASE_2)

	// extend the zeros, I could just concatinate it but idk
	bitString = fmt.Sprintf("%s%s", strings.Repeat(ZERO_APPEND, IMMEDIATE_STRING_LENGTH-len(bitString)), bitString)

	form.ImmediateValue = bitString

	return &form, nil
}
