package lettercombinationsofaphonenumber

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	digit2Letter := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	res := digit2Letter[digits[0]]
	for i := 1; i < len(digits); i++ {
		tmp := []string{}
		for j := 0; j < len(res); j++ {
			for k := 0; k < len(digit2Letter[digits[i]]); k++ {
				tmp = append(tmp, res[j]+digit2Letter[digits[i]][k])
			}
		}
		res = tmp
	}
	return res
}
