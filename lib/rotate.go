package lib

import (
	"fmt"
)

func Fw(inp Wav, t float32) Wav {
	res := inp

	fmt.Println(res._DSize)
	fmt.Println(len(res.DATA))

	fwn := int(float32(res._SRate) * t)
	cnt := int(res._DSize) / 4

	fmt.Printf("%f = %d\n", t, fwn)

	for i := 0; i < cnt; i++ {
		if i < cnt - fwn {
			res.DATA[4 * i] = res.DATA[4 * i + fwn]
			res.DATA[4 * i + 1] = res.DATA[4 * i + 1 + fwn]
			res.DATA[4 * i + 2] = res.DATA[4 * i + 2 + fwn]
			res.DATA[4 * i + 3] = res.DATA[4 * i + 3 + fwn]
		} else {
			res.DATA[4 * i] = 0
			res.DATA[4 * i + 1] = 0
			res.DATA[4 * i + 2] = 0
			res.DATA[4 * i + 3] = 0
		}
	}

	return res
}



