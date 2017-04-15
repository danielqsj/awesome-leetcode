package reverseWordsIII

import (
	"strings"
)

func reverseWordsIII(s string) string {
	var resultArray []string
	tmp := strings.Split(s, " ")
	for _, str := range tmp {
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		resultArray = append(resultArray, string(runes))
	}
	return strings.Join(resultArray, " ")
}
