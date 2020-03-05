// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"log"
	"regexp"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(s, "")
	log.Println(processedString)

	// var sigla []string
	// frase := strings.SplitAfter(processedString, " ")
	for _, s := range processedString {
		log.Println(s)
		// sigla = append(sigla, strings.ToUpper(s[:1]))
	}
	// siglaRetorno := strings.Join(sigla, "")

	// fmt.Println(processedString)
	return "algo"
}
