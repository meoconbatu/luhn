package luhn

import (
	"math"
	"strconv"
	"strings"
)

const testVersion = 2

func Valid(input string) (output bool) {
	str, err := Validate(input)
	if err {
		return false
	}
	if math.Mod(Sum(str), 10) == 0 {
		return true
	}
	return false

}
func Sum(str string) float64 {
	sum := 0
	for i := 0; i < len(str); i++ {
		if a, er := strconv.Atoi(string(str[i])); er == nil {
			if math.Mod(float64(len(str)-i), 2) == 0 {
				a = a * 2
				if a > 9 {
					a = a - 9
				}
			}
			sum = sum + a
		}
	}
	return float64(sum)
}
func Validate(input string) (output string, err bool) {
	output = strings.Replace(input, " ", "", -1)
	if len(output) <= 1 {
		return output, true
	}
	_, er := strconv.Atoi(output)
	if er != nil {
		return output, true
	}
	return output, false
}
