package domain

func getOffsetLimitTen(page int) (int, int) {
	limit := 10
	offset := limit * (page - 1)
	return limit, offset
}
