package q5

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	// Seu código aqui
	return ""
}


package main

import (
    "fmt"
    "strings"
)

func processar_sequencia(sequencia string) string {
    var resultado strings.Builder
    vogais := "aeiouAEIOU"
    for _, letra := range sequencia {
        if strings.ContainsRune(vogais, letra) {
            continue
        } else if letra >= 'A' && letra <= 'Z' {
            resultado.WriteString(strings.ToLower(string(letra)))
            resultado.WriteString(".")
        } else {
            resultado.WriteString(string(letra))
            resultado.WriteString(".")
        }
    }
    return resultado.String()
}

func main() {
    sequencia := "Hello World"
    resultado := processar_sequencia(sequencia)
    fmt.Println(resultado)
}
