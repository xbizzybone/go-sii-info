package utils

import (
	"strconv"
	"strings"
)

func IsValidRut(rut string) bool {
	rut = Format(rut)
	if len(rut) < 3 {
		return false
	}
	t := strings.Split(rut, "-")

	return VerificationCode(t[0]) == t[1]
}

func Format(rut string) string {
	rut = strings.Trim(rut, " ")
	rut = strings.TrimLeft(rut, "0")
	rut = strings.ToUpper(rut)
	rutFormated := ""
	for i := 0; i < len(rut); i++ {
		if i == len(rut)-1 {
			if isNumeric(rut[i]) {
				rutFormated += "-" + string(rut[i])
			}
			if int(rut[i]) == 75 || int(rut[i]) == 107 {
				rutFormated += "-K"
			}
		} else if isNumeric(rut[i]) {
			rutFormated += string(rut[i])
		}
	}
	return rutFormated
}

func VerificationCode(rut string) (dv string) {
	rut = strings.Trim(rut, " ")
	rut = strings.TrimLeft(rut, "0")
	factor := []int{2, 3, 4, 5, 6, 7}
	f := 0
	sum := 0
	for i := len(rut) - 1; i > -1; i-- {
		if isNumeric(rut[i]) {
			num, _ := strconv.ParseInt(string(rut[i]), 0, 64)
			sum += int(num) * factor[f]
			if f == len(factor)-1 {
				f = 0
			} else {
				f++
			}
		}
	}
	num := 11 - (sum % 11)
	if num < 10 {
		dv = strconv.Itoa(num)
	} else if num == 10 {
		dv = "K"
	} else {
		dv = "0"
	}
	return dv
}

func isNumeric(s byte) bool {
	return (int(s) >= 48 && int(s) <= 57)
}

func GetVerificationCode(rut string) string {
	return rut[len(rut)-1:]
}

func GetRutWithoutVerificationCode(rut string) string {
	return rut[:len(rut)-2]
}

func GetBooleanFromYesNoString(s string) bool {
	return s == "SI"
}

func ExtractLastCharacters(str string, length int) string {
	return strings.TrimSpace(str[len(str)-length:])
}

func FormatCategory(category string) int {
	category = strings.ToUpper(category)

	switch category {
	case "PRIMERA":
		return 1
	case "SEGUNDA":
		return 2
	default:
		return 0
	}
}
