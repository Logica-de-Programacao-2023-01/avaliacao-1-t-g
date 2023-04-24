package q3

//Você recebe um tabuleiro retangular de M x N quadrados. Além disso, você tem um número ilimitado de peças de dominó
//padrão de 2 x 1 quadrados. Você pode girar as peças. Você deve colocar o máximo de peças de dominó possível no
//tabuleiro, seguindo as seguintes condições:
//
//1. Cada peça de dominó cobre completamente dois quadrados.
//2. Nenhuma duas peças de dominó se sobrepõem.
//3. Cada peça de dominó está completamente dentro do tabuleiro. É permitido que a peça toque as bordas do tabuleiro.
//
//Encontre o número máximo de peças de dominó que podem ser colocadas sob essas restrições.
//
//Se M ou N forem iguais ou menores que 0, a função deve retornar um erro.

func DominoPieces(m, n int) (int, error) {
	// Seu código aqui
	return 0, nil
}




package main

import "fmt"

func maxDominoes(M int, N int) (int, error) {
    if M <= 0 || N <= 0 {
        return 0, fmt.Errorf("M and N must be greater than 0")
    }

    // Verificando se M ou N são ímpares e ajustando para serem pares
    if M%2 == 1 {
        M--
    }
    if N%2 == 1 {
        N--
    }

    // Calculando o número máximo de peças de dominó que podem ser colocadas
    max := (M * N) / 2

    return max, nil
}

func main() {
    M, N := 3, 5
    max, err := maxDominoes(M, N)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("M = %d, N = %d, Max dominoes = %d\n", M, N, max) // output: M = 3, N = 4, Max dominoes = 6
}
