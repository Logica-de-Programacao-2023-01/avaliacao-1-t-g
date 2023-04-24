package main

import  "errors"

func ClassifyPrices(prices []int) (int, error) {
	n := len(prices)
	if n == 0 {
		return 0, errors.New("lISTA VAZIA")
	} else if n == 1 {
		return 3, nil
	}

	cres := true
	dre := true
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			dre = false
		} else if prices[i] < prices[i-1] {
			cres = false
		}
	}

	if cres {
		return 1, nil
	} else if dre {
		return 2, nil
	} else {
		return 3, nil
	}
}
