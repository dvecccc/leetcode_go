package main

import "fmt"

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return nil
	}
	str := switchNumber(digits[0])
	for i := 1; i < n; i++ {
		var con = make([]string, 0)
		tmp := switchNumber(digits[i])
		for j := 0; j < len(tmp); j++ {
			for k := 0; k < len(str); k++ {
				con = append(con, str[k] + tmp[j])
				//fmt.Println(con)
			}
		}
		str = con[:]
	}
	return str
}

func switchNumber(str byte) []string {
	switch str {
	case '2':
		return []string{"a", "b", "c"}
	case '3':
		return []string{"d", "e", "f"}
	case '4':
		return []string{"g", "h", "i"}
	case '5':
		return []string{"j", "k", "l"}
	case '6':
		return []string{"m", "n", "o"}
	case '7':
		return []string{"p", "q", "r", "s"}
	case '8':
		return []string{"t", "u", "v"}
	case '9':
		return []string{"w", "x", "y", "z"}
	default:
		return nil
	}
}

func main()  {
	res := letterCombinations("2379")
	fmt.Println("len:", len(res))
	fmt.Println(res)
}