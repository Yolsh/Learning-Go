package worksheets

import (
	"fmt"
	"math"
)

func W110() {
	var Ans string
	fmt.Println("Give me a 2s compliment binary number")
	fmt.Scan(&Ans)
	fmt.Printf("Denary: %v", CompConvert(Ans))
}

func CompConvert(BinNum string) float32 {
	var num float32 = 0
	if BinNum[0] == '1' {
		num = -128
	}
	for i := 1; i < len(BinNum); i++ {
		if BinNum[i] == '1' {
			num += float32(math.Pow(2, float64(-i+7)))
		}
	}
	return num
}
