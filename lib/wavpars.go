package lib

import (
	"os"
	"fmt"
	"strconv"
	"encoding/binary"
)

func Wavout(fn string, inp Wav) {
	file, err := os.Create(fn)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	file.Write(inp.RHead)
	file.Write(inp.Size)
	file.Write(inp.WHead)
	file.Write(inp.FHead)

	file.Write(inp.Bytes)
	file.Write(inp.FmtId)
	file.Write(inp.Chnum)
	file.Write(inp.SRate)
	file.Write(inp.Speed)
	file.Write(inp.BSize)
	file.Write(inp.BRate)
	file.Write(inp.DHead)
	file.Write(inp.DSize)

	file.Write(inp.DATA)
}

func QBDatadumpFde(inp Wav, fn string, tbeg, tlim float32) {
	file, err := os.Create(fn)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	num := int(float32(inp._SRate) * tlim)
	bof := int(float32(inp._SRate) * tbeg)
	rd, ld := Datapars16(inp.DATA)

	if len(rd) < bof + 2 * num {
		bof = len(rd) - 2 * num
	}

	for i := 0; i < num / 4; i++ {
		file.WriteString(strconv.Itoa(i))
		file.WriteString(" ")
		sfv := fmt.Sprintf("%f", float32(rd[bof + i * 4 + num]) / float32(0x7FFF))
		file.WriteString(sfv)
		file.WriteString(" ")
		sfv = fmt.Sprintf("%f", float32(ld[bof + i * 4 + num]) / float32(0x7FFF))
		file.WriteString(sfv)
		file.WriteString("\n")
	}
}


func QDatadumpFde(inp Wav, fn string, tbeg, tlim float32) {
	file, err := os.Create(fn)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	num := int(float32(inp._SRate) * tlim)
	bof := int(float32(inp._SRate) * tbeg)
	rd, ld := Datapars16(inp.DATA)

	if len(rd) < bof + num {
		bof = len(rd) - num
	}

	for i := 0; i < num / 4; i++ {
		file.WriteString(strconv.Itoa(i))
		file.WriteString(" ")
		sfv := fmt.Sprintf("%f", float32(rd[bof + i * 4]) / float32(0x7FFF))
		file.WriteString(sfv)
		file.WriteString(" ")
		sfv = fmt.Sprintf("%f", float32(ld[bof + i * 4]) / float32(0x7FFF))
		file.WriteString(sfv)
		file.WriteString("\n")
	}
}

func DatadumpFde(inp Wav, fn string, tbeg, tlim float32) {
	file, err := os.Create(fn)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	num := int(float32(inp._SRate) * tlim)
	bof := int(float32(inp._SRate) * tbeg)
	rd, ld := Datapars16(inp.DATA)

	if len(rd) < bof + num {
		bof = len(rd) - num
	}

	for i := 0; i < num; i++ {
		file.WriteString(strconv.Itoa(i))
		file.WriteString(" ")
		sfv := fmt.Sprintf("%f", float32(rd[bof + i]) / float32(0x7FFF))
		file.WriteString(sfv)
		file.WriteString(" ")
		sfv = fmt.Sprintf("%f", float32(ld[bof + i]) / float32(0x7FFF))
		file.WriteString(sfv)
		file.WriteString("\n")
	}
}

func Datadump16de(inp Wav, fn string, tbeg, tlim float32) {
	file, err := os.Create(fn)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	num := int(float32(inp._SRate) * tlim)
	bof := int(float32(inp._SRate) * tbeg)
	rd, ld := Datapars16(inp.DATA)

	if len(rd) < bof + num {
		bof = len(rd) - num
	}

	for i := 0; i < num; i++ {
		file.WriteString(strconv.Itoa(i))
		file.WriteString(" ")
		file.WriteString(strconv.Itoa(int(rd[bof + i])))
		file.WriteString(" ")
		file.WriteString(strconv.Itoa(int(ld[bof + i])))
		file.WriteString("\n")
	}
}

func Datadump16(inp Wav, fn string, tlim float32) {
	file, err := os.Create(fn)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	num := int(float32(inp._SRate) * tlim)
	rd, ld := Datapars16(inp.DATA)

	for i := 0; i < num; i++ {
		file.WriteString(strconv.Itoa(i))
		file.WriteString(" ")
		file.WriteString(strconv.Itoa(int(rd[i])))
		file.WriteString(" ")
		file.WriteString(strconv.Itoa(int(ld[i])))
		file.WriteString("\n")
	}
}

func Dataparsf64(inp []byte) ([]float64, []float64) {
	rd := make([]float64, len(inp) / 4)
	ld := make([]float64, len(inp) / 4)

	for i := 0; i < len(inp) / 4; i++ {
		cache := int16(inp[4 * i + 1]) * int16(0x100) + int16(inp[4 * i + 0])
		rd[i] = float64(cache) / float64(0x8000)
		cache = int16(inp[4 * i + 3]) * int16(0x100) + int16(inp[4 * i + 2])
		ld[i] = float64(cache) / float64(0x8000)
	}

	return rd, ld
}

func Wavpars(fn string) Wav {
	file, err := os.Open(fn)
	defer file.Close()

	var res Wav

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fn, "open")
	}

	res.RHead = make([]byte, 4)
	res.Size = make([]byte, 4)
	res.WHead = make([]byte, 4)
	res.FHead = make([]byte, 4)

	res.Bytes = make([]byte, 4)
	res.FmtId = make([]byte, 2)
	res.Chnum = make([]byte, 2)
	res.SRate = make([]byte, 4)
	res.Speed = make([]byte, 4)
	res.BSize = make([]byte, 2)
	res.BRate = make([]byte, 2)
	res.DHead = make([]byte, 1)
	res.DSize = make([]byte, 4)

	file.Read(res.RHead)
	file.Read(res.Size)
	res._Size = binary.LittleEndian.Uint32(res.Size)
	file.Read(res.WHead)
	file.Read(res.FHead)

	file.Read(res.Bytes)
	file.Read(res.FmtId)
	file.Read(res.Chnum)
	file.Read(res.SRate)
	res._SRate = binary.LittleEndian.Uint32(res.SRate)
	file.Read(res.Speed)
	file.Read(res.BSize)
	file.Read(res.BRate)

	file.Read(res.DHead)
	for {
		for res.DHead[0] != 0x64 {
			file.Read(res.DHead)
		}
		rata := make([]byte, 3)
		sata := [] byte{0x61, 0x74, 0x61}
		file.Read(rata)
		jata := 0
		for i := 0; i < 3; i ++ {
			jata += int(rata[i] - sata[i])
		}
		if jata == 0{
			res.DHead = append(res.DHead, 'a', 't', 'a')
			break
		}
	}

	file.Read(res.DSize)
	res._DSize = binary.LittleEndian.Uint32(res.DSize)

	res.DATA = make([]byte, res._DSize)
	file.Read(res.DATA)

	fmt.Printf("%s\n", res.RHead)
	fmt.Println(res._Size)
	fmt.Printf("%s\n", res.WHead)
	fmt.Printf("%s\n", res.FHead)
	fmt.Println(res.Bytes)
	fmt.Println(res._SRate)
	fmt.Println(res.BRate)
	fmt.Printf("%s\n", res.DHead)
	fmt.Println(res._DSize)

	return res
}
