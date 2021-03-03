package hammingdistance

func hammingDistance(x int, y int) int {
	x ^= y
	count := 0
	for x > 0 {
		x &= (x - 1)
		count++
	}
	return count
}
