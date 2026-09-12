package models

type ISA map[string]map[string]string // bad romance
type Registers map[string]string
type Instruction interface {
	ConvertToBinary()
}
