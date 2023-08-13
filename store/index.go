package store

import (
	"strings"
)

func isin(a int, b []int) bool {
	for _, v := range b {
		if v == a {
			return true
		}
	}
	return false
}

func Tokenize(phrases []string) map[string][]int {
	words := map[string][]int{}
	for i := 0; i < len(phrases); i++ {
		sentence := strings.Fields(phrases[i])
		for _, val := range sentence {
			words[val] = append(words[val], i)
		}
	}
	return words
}

func Bigram(filenames []string) map[string][]int {
	words := map[string][]int{}
	for i := 0; i < len(filenames); i++ {
		name := []rune(filenames[i])
		for j := 0; j < len(name)-1; j++ {
			bi := string(name[j : j+2])
			if !isin(i, words[bi]) {
				words[bi] = append(words[bi], i)
			}
		}
	}
	return words
}

func Unigram(filenames []string) map[string][]int {
	words := map[string][]int{}
	for i := 0; i < len(filenames); i++ {
		name := []rune(filenames[i])
		for j := 0; j < len(name); j++ {
			uni := string(name[j])
			words[uni] = append(words[uni], i)
		}
	}
	return words
}
