package find

import (
	"sort"
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

func Rank(find string, found_indicies []int, bigrams map[string][]int, names []string) []int {
	first := string([]rune(find)[0:2])

	found := []string{}
	for _, val := range found_indicies {
		found = append(found, names[val])
	}

	from_start := []int{} // use percentage instead of absolute index. example, word could be sort, but letter searched for is at the end
	for _, val := range found {
		characters := []rune(val)
		for i := 0; i < len(val)-1; i++ {
			if string(characters[i:i+2]) == first {
				from_start = append(from_start, i)
				break
			}
		}
	}

	// delete this and do it yourself
	sortedIndices := make([]int, len(found_indicies))
	for i := range found_indicies {
		sortedIndices[i] = i
	}

	sort.SliceStable(sortedIndices, func(i, j int) bool {
		return from_start[sortedIndices[i]] < from_start[sortedIndices[j]]
	})

	result := make([]int, len(found_indicies))
	for i, idx := range sortedIndices {
		result[i] = found_indicies[idx]
	}
	// delete this and do it yourself

	return result
}

func Search(find string, unigrams map[string][]int, bigrams map[string][]int) []int {
	if len(find) == 1 {
		return unigrams[find]
	}
	characters := []rune(find)
	least := bigrams[string(characters[0:2])]
	for i := 0; i < len(find)-1; i++ {
		found := bigrams[string(characters[i:i+2])]
		if len(found) < len(least) {
			least = found
		} else if len(found) == len(least) {
			temp := []int{}
			for _, v := range found {
				if isin(v, least) {
					temp = append(temp, v)
				}
			}
			least = temp
		}
	}
	return least
}

func DistanceRank(find string, found_indicies []int, bigrams map[string][]int, names []string) ([]int, []int) {
	// get 'find' bigrams
	characters := []rune(find)
	used_bigrams := map[string]bool{}
	for i := 0; i < len(find)-1; i++ {
		bi := string(characters[i : i+2])
		if _, ok := bigrams[bi]; ok {
			used_bigrams[bi] = true
		}
	}

	found := []string{}
	for _, val := range found_indicies {
		found = append(found, names[val])
	}

	fs, ad := []int{}, []int{}
	if len(found) > 3 {
		found = found[0:3]
	}
	for _, val := range found {
		characters := []rune(val)
		from_start := 0
		avg_distance := 0
		last_bigram := 0
		for j := 0; j < len(val)-1; j++ {
			if _, ok := used_bigrams[string(characters[j:j+2])]; ok {
				if from_start == 0 && last_bigram == 0 {
					from_start = j
				}
				avg_distance += j + 1 - last_bigram
				last_bigram = j + 1
			}
		}
		avg_distance /= len(val)
		fs = append(fs, from_start)
		ad = append(ad, avg_distance)

	}
	return fs, ad
}

func SentenceSearch(find string, words map[string][]int) []int {
	sentence := strings.Fields(find)
	least := words[sentence[0]]
	for _, val := range sentence {
		found := words[val]
		if len(found) < len(least) {
			least = found
		} else if len(found) == len(least) {
			temp := []int{}
			for _, v := range found {
				if isin(v, least) {
					temp = append(temp, v)
				}
			}
			least = temp
		}
	}
	return least
}
