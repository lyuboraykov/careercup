package reversewords

import "strings"

// https://www.careercup.com/question?id=5697358784364544

func reverseWordsInSentence(sentence string) string {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		words[i] = reverseString(word)
	}
	return strings.Join(words, " ")
}

func reverseString(s string) string {
	res := make([]rune, 0)
	for _, r := range s {
		res = append([]rune{r}, res...)
	}
	return string(res)
}
