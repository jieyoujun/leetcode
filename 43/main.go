package multiplystrings

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	resInts := make([]int, len(num1)+len(num2))
	for in2 := len(num2) - 1; in2 >= 0; in2-- {
		n2 := int(num2[in2] - '0')
		for in1 := len(num1) - 1; in1 >= 0; in1-- {
			n1 := int(num1[in1] - '0')
			sum := n2*n1 + resInts[in2+in1+1]
			resInts[in2+in1+1] = sum % 10 // 存进位
			resInts[in2+in1] += sum / 10
		}
	}
	if resInts[0] == 0 {
		resInts = resInts[1:]
	}
	resStr := ""
	for _, r := range resInts {
		resStr += string(rune(r + '0'))
	}
	return resStr
}
