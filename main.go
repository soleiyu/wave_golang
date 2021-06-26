package main

import (
	"./lib"
	"os"
	"fmt"
)

func main() {
	println("wavart")

	wav := lib.Wavpars("ss.wav")

	rd, _ := lib.Datapars16(wav.DATA)
	mid := wav.SRate_
	dump(rd[mid - mid / 5: mid + mid / 5])

	an, bn := lib.WindowDft(rd[mid - mid / 5: mid + mid / 5])
	lib.ABndump(an, bn)
//	lib.Andump(an)
//	lib.Bndump(bn)

//	lib.Wavout("res.wav", wav)
}

func dump(inp []int16) {

	file, _ := os.Create("dump")
	defer file.Close()

	for i := 0; i < len(inp); i++ {
		sfv := fmt.Sprintf("%d, %d \n", i, inp[i])
		file.WriteString(sfv)
	}
}
