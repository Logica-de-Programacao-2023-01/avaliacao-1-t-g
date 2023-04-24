package q2

func ProblemsSolved(answers [][3]bool) int {
	count := 0
	for _, opiniao := range answers {
		sum := 0
		for _, hasSolution := range opiniao {
			if hasSolution {
				sum++
			}
		}
		if sum >= 2 {
			count++
		}
	}
	return count
}

