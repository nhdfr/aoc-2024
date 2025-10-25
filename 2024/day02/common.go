package main

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}
	increasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if increasing && diff <= 0 || !increasing && diff >= 0 {
			return false
		}
		absDiff := diff
		if absDiff < 0 {
			absDiff = -absDiff
		}
		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}
	return true
}
