package pendu

import (
    "fmt"
    "strings"
)

func ValidLetter(finalWord, hiddenWord, letter string) (string, int) {
    if strings.Contains(finalWord, letter) {
        fmt.Println("La lettre choisie est dans le mot")

        // Convertir hiddenWord et finalWord en slices de runes pour pouvoir modifier hiddenWord
        hiddenRunes := []rune(hiddenWord)
        finalRunes := []rune(finalWord)

        for i := 0; i < len(finalRunes); i++ {
            if finalRunes[i] == rune(letter[0]) && hiddenRunes[i] == '.' {
                hiddenRunes[i] = rune(letter[0])
            }
        }

        // Retourner hiddenWord mis à jour en chaîne
        return string(hiddenRunes), 0
    }

    fmt.Println("La lettre n'est pas présente dans le mot")
    return hiddenWord, 1
}
