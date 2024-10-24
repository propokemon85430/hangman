package github.com/propokemon85430/hangman

import (
	"math/rand/v2"
	"os"
	"strings"
)

func Random_Word() string {
	tab := [3]string{"words.txt", "words2.txt", "words3.txt"}
	index := rand.IntN(len(tab))
	filePath := "../" + tab[index]
	content, error := os.ReadFile(filePath)
	if error != nil {
		return "File name missing"
	}
	lines := strings.Split(string(content), "\n")
	random_word := rand.IntN(22)
	new_word := (string(lines[random_word]) + "\n")
	return new_word
}
