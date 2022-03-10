package main

import (
	"fmt"
	"math"
)



func intToRoman(num int) string {
	var count int
	var strTemp = make([]string, 0)
	var tempStr string

	for num > 0 {
		tempStr = ""
		temp := num % 10
		num = num / 10
		realValue := temp * int(math.Pow10(count))

		if v := covToRoman(realValue); v != "" {
			strTemp = append(strTemp, v)
		} else {
			if temp-5 > 0 {
				tempStr = tempStr + covToRoman(int(5 * math.Pow10(count)))
				ssr := covToRoman(int(math.Pow10(count)))
				for temp-5 > 0 {
					tempStr = tempStr + ssr
					temp--
				}
				//fmt.Println(covToRoman(int(5 * math.Pow10(count))))
			} else {
				ssr := covToRoman(int(math.Pow10(count)))
				for temp > 0 {
					tempStr = tempStr + ssr
					temp--
				}
			}
			strTemp = append(strTemp, tempStr)
		}
		count++
		//println(count)
	}
	var s string
	for index := len(strTemp) - 1; index >= 0; index-- {
		s += strTemp[index]
	}
	return s
}

func covToRoman(num int) string {
	switch num {
	case 1:
		return "I"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 9:
		return "IX"
	case 10:
		return "X"
	case 40:
		return "XL"
	case 50:
		return "L"
	case 90:
		return "XC"
	case 100:
		return "C"
	case 400:
		return "CD"
	case 500:
		return "D"
	case 900:
		return "CM"
	case 1000:
		return "M"
	default:
		return ""
	}
}

func main()  {
	str := intToRoman(3999)
	fmt.Println(str)
}
