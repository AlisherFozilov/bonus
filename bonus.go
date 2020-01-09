package bonus

func bonus(sales []int) int {
	const percent = 5
	const bound = 10_000
	result := 0

	for _, amount := range sales {
		if amount > bound {
			extra := amount - bound
			result += extra * percent / 100
		}
	}

	return result
}
