package pendu

import (
    "fmt"
    "strings"
)

func ValidLetter(finalWord, hiddenWord, letter string) (string, int) {
    if strings.Contains(finalWord, letter) {
        fmt.Println("La lettre choisie est dans le mot")

        for i := 0; i < len(finalWord); i++ {
            if finalWord[i] == letter[0]) && hiddenWord[i] == '.' {
                hiddenWord[i] = letter[0]
            }
        }

        return hiddenWord, 0
    }

    fmt.Println("La lettre n'est pas présente dans le mot")
    return hiddenWord, 1
}
