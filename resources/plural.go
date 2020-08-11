package resources

func intInRange(i, from, to int64) bool {
	return from <= i && i <= to
}

func intEqualsAny(i int64, any ...int64) bool {
	for _, a := range any {
		if i == a {
			return true
		}
	}
	return false
}
