package couplesholdinghands

func minSwapsCouples(row []int) int {
	n := len(row)
	person2Position := make([]int, n)
	for position, person := range row {
		person2Position[person] = position
	}
	swap := 0
	for i := 0; i < n; i += 2 {
		loverOfRowI := row[i] ^ 1
		if loverOfRowI == row[i+1] {
			continue
		}
		wrongLover := row[i+1]
		row[i+1], row[person2Position[loverOfRowI]] = row[person2Position[loverOfRowI]], row[i+1]
		person2Position[loverOfRowI], person2Position[wrongLover] = person2Position[wrongLover], person2Position[loverOfRowI]
		swap++
	}
	return swap
}
