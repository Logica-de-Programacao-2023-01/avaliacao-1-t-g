package q5

import 	"strings"

func ProcessString(s string) string {
	var resultado strings.Builder
	vogais := "aeiouAEIOU"
	for _, letra := range s {
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

