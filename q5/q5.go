package q5

import 	"strings"

func ProcessString(s string) string {
	var resultado strings.Builder
	vogais := "aeiouAEIOU"
	for _, letra := range s {
		if strings.ContainsRune(vogais, letra) {
			continue
		} else if letra >= 'A' && letra <= 'Z' {
			resultado.WriteString(".")
			resultado.WriteString(strings.ToLower(string(letra)))
			
		} else {
			resultado.WriteString(".")
			resultado.WriteString(string(letra))
			
		}
	}
	return resultado.String()
}

