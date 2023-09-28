package roman_to_integer

func romanToInt(s string) int {
	numerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	num := 0
	prevRoman := ""

	for i := len(s) - 1; i >= 0; i-- {
		roman := string(s[i])
		if roman == "I" && (prevRoman == "V" || prevRoman == "X") {
			num -= 1
			prevRoman = roman
			continue
		}

		if roman == "X" && (prevRoman == "C" || prevRoman == "L") {
			num -= 10
			prevRoman = roman
			continue
		}

		if roman == "C" && (prevRoman == "M" || prevRoman == "D") {
			num -= 100
			prevRoman = roman
			continue
		}

		prevRoman = roman
		num += numerals[roman]
	}

	return num
}
