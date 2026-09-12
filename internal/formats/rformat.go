package formats

type RFromat struct {
	Instruction         string
	FirstRegister       string
	SecondRegister      string
	DestinationRegister string
}

func (rf *RFromat) ConvertToBinary() string {
	binary := ""

	return binary
}
