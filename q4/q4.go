package q4

//Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
//deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
//lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.
//
//Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
//ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
//estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
//Caso a lista possua apenas um elemento, a função deve retornar 3.

func ClassifyPrices(prices []int) (int, error) {
	// Seu código aqui
	return 0, nil
}


package main

import (
    "errors"
    "fmt"
)

func classifyList(prices []int) (int, error) {
    n := len(prices)
    if n == 0 {
        return 0, errors.New("empty list")
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

func main() {
    prices := []int{1, 2, 3, 4, 5}
    result, err := classifyList(prices)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result) // output: 1
}
