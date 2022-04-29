package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("-91283472332"))
}

func myAtoi(s string) int {
	str := []rune(s)
	strLength := len(str)
	i := 0
	for i < strLength && str[i] == ' ' {
		i++
	}

	flag := 0

	if i < strLength && str[i] == '-' {
		flag = -1
		i++
	}

	if flag == 0 && i < strLength {
		flag = +1
		if str[i] == '+' {
			i++
		}
	}

	result := 0
	for i < strLength && str[i] >= '0' && str[i] <= '9' {
		digit := int(str[i] - '0')
		if flag == 1 && result > (math.MaxInt32-digit)/10 {
			return math.MaxInt32
		}
		if flag == -1 && result > (math.MaxInt32-digit)/10 {
			return math.MinInt32
		}
		result = result*10 + digit
		i++
	}
	return result * flag
}
