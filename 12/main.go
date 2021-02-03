package integertoroman

func IntToRoman(num int) string {
	var roman string
	if num < 0 || num > 4000 {
		return roman
	}
	intParts := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanParts := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(intParts); i++ {
		for num >= intParts[i] {
			num -= intParts[i]
			roman += romanParts[i]
		}
	}
	return roman
}
