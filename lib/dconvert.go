package lib

func BRChange(ld, rd []int16, obr, nbr int) ([]int16, []int16) {
	newcount := len(rd) * nbr / obr

	resr := make([]int16, newcount)
	resl := make([]int16, newcount)

	for i := 0; i < newcount; i++ {
		resr[i] = rd[i * obr / nbr]
		resl[i] = ld[i * obr / nbr]
	}

	return resr, resl
}

func Datamarg16(ld, rd []int16) []byte {
	res := make([]byte, len(ld) * 4)

	for i := 0; i < len(ld); i++ {
		res[4 * i + 1] = byte(rd[i] >> 8)
		res[4 * i + 0] = byte(rd[i] % int16(0x100))
		res[4 * i + 3] = byte(ld[i] >> 8)
		res[4 * i + 2] = byte(ld[i] % int16(0x100))
	}

	return res
}

func Datapars8l(inp []byte) ([]int16, []int16) {
	rd := make([]int16, len(inp) / 4)
	ld := make([]int16, len(inp) / 4)

	for i := 0; i < len(inp) / 4; i++ {
		rd[i] = int16(inp[4 * i + 1]) << 8
		ld[i] = int16(inp[4 * i + 3]) << 8
	}

	return rd, ld
}

func Datapars8r(inp []byte) ([]byte, []byte) {
	rd := make([]byte, len(inp) / 4)
	ld := make([]byte, len(inp) / 4)

	for i := 0; i < len(inp) / 4; i++ {
		rd[i] = inp[4 * i + 1]
		ld[i] = inp[4 * i + 3]
	}

	return rd, ld
}

func Datapars16(inp []byte) ([]int16, []int16) {
	rd := make([]int16, len(inp) / 4)
	ld := make([]int16, len(inp) / 4)

	for i := 0; i < len(inp) / 4; i++ {
		rd[i] = int16(inp[4 * i + 1]) * int16(0x100) + int16(inp[4 * i + 0])
		ld[i] = int16(inp[4 * i + 3]) * int16(0x100) + int16(inp[4 * i + 2])
	}

	return rd, ld
}

