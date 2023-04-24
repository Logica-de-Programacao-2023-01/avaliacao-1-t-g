package q4

func ClassifyPrices(prices []int) (int, error) {
	n := len(prices)
	if n == 0 {
		return 0, errors.New("lISTA VAZIA")
	} else if n == 1 {
		return 3, nil
	}

	isIncreasing := true
	isDecreasing := true
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			isDecreasing = false
		} else if prices[i] < prices[i-1] {
			isIncreasing = false
		}
	}

	if isIncreasing {
		return 1, nil
	} else if isDecreasing {
		return 2, nil
	} else {
		return 3, nil
	}
}
