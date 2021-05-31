package main

import (
	"./lib"
)

func main() {
	println("wavart")

	wav := lib.Wavpars("ss.wav")

	lib.Fw(wav, 0.3)

//	rd, ld := lib.Datapars16(wav.DATA)

//	r11, l11 := lib.BRChange(ld, rd, 44100, 11520)
//	r44, l44 := lib.BRChange(l11, r11, 11520, 44100)

//	rd8l, ld8l := lib.Datapars8l(wav.DATA)

//	wav.DATA = lib.Datamarg16(rd8l, ld8l)
//	wav.DATA = lib.Datamarg16(r11, l11)
//	rd8r, _:= lib.Datapars8r(wav.DATA)

	lib.Wavout("res.wav", wav)
}
