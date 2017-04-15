package keyboardRow

import (
	"strings"
)

func keyboardRow(words []string) []string {
	templates := []string{"QWERTYUIOP", "ASDFGHJKL", "ZXCVBNM"}
	tempHash := make(map[rune]int)
	for i, template := range templates {
		runes := []rune(template)
		for _, char := range runes {
			tempHash[char] = i
		}
	}

	var result []string
loop:
	for _, word := range words {
		upperWord := strings.ToUpper(word)
		runes := []rune(upperWord)
		lineNumber := tempHash[runes[0]]
		for _, char := range runes {
			if tempHash[char] != lineNumber {
				continue loop
			}
		}
		result = append(result, word)
	}
	return result
}
