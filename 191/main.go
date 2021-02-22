package numberof1bits

func hammingWeight(num uint32) int {
	// n&n-1, 去掉最后一个1
	// res := 0
	// for num != 0 {
	// 	res++
	// 	num &= (num - 1)
	// }
	// return res
	// redis swar
	num = (num & 0x55555555) + ((num >> 1) & 0x55555555)
	num = (num & 0x33333333) + ((num >> 2) & 0x33333333)
	num = (num & 0x0f0f0f0f) + ((num >> 4) & 0x0f0f0f0f)
	num = (num * (0x01010101) >> 24)
	return int(num)
}
