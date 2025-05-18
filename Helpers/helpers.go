package Helpers

import "fmt"

func FormatSalary(salary int) string {
	str := fmt.Sprintf("%d", salary)
	result := ""

	for i, digit := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += "."
		}
		result += string(digit)
	}

	return "Rp " + result
}