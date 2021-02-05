package sqrtx

func mySqrt(x int) int {
	// 如何快速定位
	// 直接遍历也太low了吧

	// 二分啊啊啊啊
	left, right := 0, x
	for left <= right {
		middle := (left + right) / 2
		if middle*middle > x {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left - 1

	// 胡思乱想
	// // 定位
	// // x=101<2^7
	// // x=8<2^4
	// a := 0
	// for i := x; i > 0; i = i >> 1 {
	// 	a++
	// }
	// fmt.Println(x, a)

	// // sqrtx in 2^3~2^4
	// // sqrtx in 2^1~2^2
	// i := 1 << (a/2 - 1)
	// for {
	// 	if (i+1)*(i+1) > x {
	// 		fmt.Println(i)
	// 		break
	// 	}
	// 	i++
	// }
	// return i
}
