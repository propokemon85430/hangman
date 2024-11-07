package pendu

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

func CreateWord(filePath string) (string, string) {
	new_word := Random_Word(filePath)
	index := 0
	hidden_word := ""
	beginning_word := ""
	var random_index []int
	for i := 0; i < len(new_word); i++ {
		hidden_word = hidden_word + "."
	}
	for l := 0; l < len(new_word)/2-1; l++ {
		random_index = append(random_index, rand.IntN(len(new_word)))
	}

	sort.Ints(random_index)

	for j := 0; j < len(hidden_word); j++ {
		if random_index[index] == j {
			beginning_word = beginning_word + string(new_word[j])
			index = index + 1
		} else {
			beginning_word = beginning_word + "."
		}
		if index == len(random_index) {
			index = 0
		}
	}
	fmt.Println(beginning_word)
	return beginning_word, new_word
}
