package choushulcof

func NthUglyNumber(n int) int {
	if n > 0 {
		ugly := make([]int, n)
		ugly[0] = 1
		i, j, k := 0, 0, 0
		for idx := 1; idx < n; idx++ {
			ugly[idx] = min(ugly[i]*2, ugly[j]*3, ugly[k]*5)
			if ugly[idx] == ugly[i]*2 {
				i++
			}
			if ugly[idx] == ugly[j]*3 {
				j++
			}
			if ugly[idx] == ugly[k]*5 {
				k++
			}
		}
		return ugly[n-1]
	}
	return 0
}
func min(n1, n2, n3 int) int {
	if n1 > n2 {
		n1 = n2
	}
	if n1 > n3 {
		n1 = n3
	}
	return n1
}
