package bonus

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

func CalculateTowers(barLengths []int) (int, int) {
	// Seu código aqui
	return 0, 0
}



func buildTowers(bars []int) (int, int) {
    // Cria um mapa para contar a frequência dos comprimentos das barras
    freq := make(map[int]int)
    for _, bar := range bars {
        freq[bar]++
    }
    
    // Encontra o comprimento mínimo das barras e a quantidade máxima de torres
    minLen := bars[0]
    maxTowers := freq[minLen]
    for len, count := range freq {
        if count > maxTowers {
            maxTowers = count
        }
        if len < minLen {
            minLen = len
        }
    }
    
    // Retorna a altura máxima das torres e a quantidade de torres
    return minLen, maxTowers
}
