package roman_number

import (
	"strings"
)

// earlier..
type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value

		}
	}
	return result.String()
}

// later..
func ConvertToArabic(roman string) int {
	total := 0
	for _, numeral := range allRomanNumerals {
		symbolLen := len(numeral.Symbol)
		for symbolLen <= len(roman) && roman[:symbolLen] == numeral.Symbol {
			total += numeral.Value
			roman = roman[symbolLen:]
		}
	}
	return total
}
