package isa

import "github.com/fingernailz/assembler/internal/models"

type Registers map[string]string

const (
	R_FORMAT  = "R-FORMAT"
	I_FORMAT  = "I-FORMAT"
	D_FORMAT  = "D-FORMAT"
	B_FORMAT  = "B-FORMAT"
	CB_FORMAT = "CB-FORMAT"
	IW_FORMAT = "IW-FORMAT"
)

var Instructions models.ISA = models.ISA{
	// R format
	"ADD": {
		"format":  R_FORMAT,
		"op-code": "10001011000",
	},
	"SUB": {
		"format":  R_FORMAT,
		"op-code": "11001011000",
	},
	"ADDS": {
		"format":  R_FORMAT,
		"op-code": "10101011000",
	},
	"SUBS": {
		"format":  R_FORMAT,
		"op-code": "11101011000",
	},
	"AND": {
		"format":  R_FORMAT,
		"op-code": "10001010000",
	},
	"ANDS": {
		"format":  R_FORMAT,
		"op-code": "11101010000",
	},
	"ORR": {
		"format":  R_FORMAT,
		"op-code": "10101010000",
	},
	"EOR": {
		"format":  R_FORMAT,
		"op-code": "11101010000",
	},
	"LSL": {
		"format":  R_FORMAT,
		"op-code": "11010011011",
	},
	"LSR": {
		"format":  R_FORMAT,
		"op-code": "11010011010",
	},
	"BR": {
		"format":  R_FORMAT,
		"op-code": "11010110000",
	},

	// I format
	"ADDI": {
		"format":  I_FORMAT,
		"op-code": "1001000100",
	},
	"SUBI": {
		"format":  I_FORMAT,
		"op-code": "1101000100",
	},
	"ADDIS": {
		"format":  I_FORMAT,
		"op-code": "1011000100",
	},
	"SUBIS": {
		"format":  I_FORMAT,
		"op-code": "1111000100",
	},
	"ANDI": {
		"format":  I_FORMAT,
		"op-code": "1001001000",
	},
	"ANDIS": {
		"format":  I_FORMAT,
		"op-code": "1111001000",
	},
	"ORRI": {
		"format":  I_FORMAT,
		"op-code": "1011001000",
	},
	"EORI": {
		"format":  I_FORMAT,
		"op-code": "1101001000",
	},

	// D format
	"LDUR": {
		"format":  D_FORMAT,
		"op-code": "11111100010",
	},
	"STUR": {
		"format":  D_FORMAT,
		"op-code": "11111100000",
	},
	"LDURSW": {
		"format":  D_FORMAT,
		"op-code": "10111000100",
	},
	"STURW": {
		"format":  D_FORMAT,
		"op-code": "10111000000",
	},
	"LDURH": {
		"format":  D_FORMAT,
		"op-code": "01111000010",
	},
	"STURH": {
		"format":  D_FORMAT,
		"op-code": "01111000000",
	},
	"LDURB": {
		"format":  D_FORMAT,
		"op-code": "00111000010",
	},
	"STURB": {
		"format":  D_FORMAT,
		"op-code": "00111000000",
	},
	"LDXR": {
		"format":  D_FORMAT,
		"op-code": "11001000010",
	},
	"STXR": {
		"format":  D_FORMAT,
		"op-code": "11001000000",
	},

	// IW format
	"MOVZ": {
		"format":  IW_FORMAT,
		"op-code": "110100101",
	},
	"MOVK": {
		"format":  IW_FORMAT,
		"op-code": "111100101",
	},

	// CB format

	"CBZ": {
		"format":  CB_FORMAT,
		"op-code": "10110100",
	},
	"CBNZ": {
		"format":  CB_FORMAT,
		"op-code": "10110101",
	},
	// branch with condition, idk how to do, do later! opcode 01010100

	// B format
	"B": {
		"format":  B_FORMAT,
		"op-code": "000101",
	},
	"BL": {
		"format":  B_FORMAT,
		"op-code": "100101",
	},
}

/*
R format
	1. 3 operands
	2. 2 operands, one shmt
	3. BR with one operand mostly x30
*/

/* D format, only prob is parsing the brackets */

/* B has no prob afaics sw CB ig*/

/* two pass assemblers so first find labels and their addresses*/
