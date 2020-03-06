// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"log"
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

	// reemplazar el espacio por la siguiente letra en mayuscula
	// convertir la frase en un array, por cada elemento evaluar por la expresion regular para reemplazo de elementos que no sean letras
	processedString1 := strings.ReplaceAll(s, " ", "")
	log.Println("processedString1: ", processedString1)

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString2 := reg.ReplaceAllString(processedString1, "")
	log.Println("processedString2: ", processedString2)

	// processedString3 := []rune(processedString2)
	campos := strings.Split(processedString2, "")
	// log.Println("processedString3: ", processedString3)

	j := 0
	for range campos {
		// runa := []rune(s)
		log.Println(campos[j])
		j++
	}

	// var sigla []string
	// // // frase := strings.SplitAfter(processedString, " ")
	// for _, s := range processedString3 {
	// 	// log.Println("valor de s: ", s)
	// 	// sigla = append(sigla, strings.ToUpper(s[:1]))
	// 	runa := []rune(s)
	// 	if unicode.IsUpper(runa[0]) {
	// 		// log.Println("mayuscula: ", s)
	// 		sigla = append(sigla, s)
	// 	}
	// }
	// siglaRetorno := strings.Join(sigla, "")
	// // log.Println("siglaRetorno: ", siglaRetorno)
	// // // fmt.Println(processedString)
	// return siglaRetorno
	return "siglaRetorno"
}
