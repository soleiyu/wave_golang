package main

import (
	"./lib"
	"strconv"
	"os"
)

func main() {
	println("wavart")
	println(os.Args[1])

	wav := lib.Wavpars(os.Args[1])
	t, _ := strconv.ParseFloat(os.Args[2], 32)

	res := lib.Fw(wav, float32(t))

	lib.Wavout("res.wav", res)
}
