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
	new_word := strings.TrimSpace(lines[random_word])
	return new_word
}
