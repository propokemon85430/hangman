package pendu

import (
	"math/rand/v2"
	"os"
	"strings"
)

func Random_Word(filePath string) string {
	content, error := os.ReadFile(filePath)
	if error != nil {
		return "File name missing"
	}
	lines := strings.Split(string(content), "\n")
	random_word := rand.IntN(22)
	new_word := (string(lines[random_word]) + "\n")
	return new_word
}
