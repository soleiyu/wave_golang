package lib

import (
	"math"
	"os"
	"fmt"
)

func Dft2(inp Wav) {
	println("dft")
	println(inp.SRate_)
}

func WindowDftA1(inp []int16) {
	println(len(inp))
	T := float64(len(inp))

	a1 := 0.0
	for t := 0; t < len(inp); t++ {
		a1 += float64(inp[t]) * math.Cos(
			2.0 * math.Pi * 1.0 * float64(t) / T)
	}
	a1 = a1 * 2.0 / T
	fmt.Printf("%f \n", a1)
}

func WindowDft(inp []int16) ([]float64,[]float64) {
	num := 1000

	return WindowDftAn(inp, num), WindowDftBn(inp, num)
}

func WindowDftAn(inp []int16, num int) []float64 {
	println(len(inp))
	T := float64(len(inp))

	an := make([]float64, num)

	for n := 0; n < num; n ++ {
		an[n] = 0.0
		for t := 0; t < len(inp); t++ {
			an[n] += float64(inp[t]) * math.Cos(
				2.0 * math.Pi * float64(n + 1) * float64(t) / T)
		}
		an[n] = an[n] * 2.0 / T
//		fmt.Printf("%f \n", an[n]);
	}

	return an
}

func WindowDftBn(inp []int16, num int) []float64 {
	println(len(inp))
	T := float64(len(inp))

	an := make([]float64, num)

	for n := 0; n < num; n ++ {
		an[n] = 0.0
		for t := 0; t < len(inp); t++ {
			an[n] += float64(inp[t]) * math.Sin(
				2.0 * math.Pi * float64(n + 1) * float64(t) / T)
		}
		an[n] = an[n] * 2.0 / T
//		fmt.Printf("%f \n", an[n]);
	}

	return an
}

func ABndump(an, bn []float64) {

	file, _ := os.Create("ABndump")
	defer file.Close()

	for i := 0; i < len(an); i++ {
		sfv := fmt.Sprintf("%d, %f, %f, %f \n", i, an[i] + bn[i], an[i], bn[i])
		file.WriteString(sfv)
	}
}

func Andump(inp []float64) {

	file, _ := os.Create("Andump")
	defer file.Close()

	for i := 0; i < len(inp); i++ {
		sfv := fmt.Sprintf("%d, %f \n", i, inp[i])
		file.WriteString(sfv)
	}
}

func Bndump(inp []float64) {

	file, _ := os.Create("Bndump")
	defer file.Close()

	for i := 0; i < len(inp); i++ {
		sfv := fmt.Sprintf("%d, %f \n", i, inp[i])
		file.WriteString(sfv)
	}
}
