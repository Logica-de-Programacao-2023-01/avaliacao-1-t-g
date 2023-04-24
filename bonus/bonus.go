package bonus

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

func CalculateTowers(barLengths []int) (int, int) {
	minLen := barLengths[0]
	maxTo := 0

	counts := make([]int, len(barLengths))
	for _, bar := range barLengths {
		counts[bar-minLen]++
		if counts[bar-minLen] > maxTo {
			maxTo = counts[bar-minLen]
		}
	}

	numTowers := maxTo
	for _, count := range counts {
		if count > 0 {
			numTowers++
		}
	}

	return minLen, barLengths[2]
}
