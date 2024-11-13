package pendu

import (
    "fmt"
    "strings"
)

func ValidLetter(finalWord, hiddenWord, letter string) (string, int) {

    if strings.Contains(finalWord, letter) {
        fmt.Println("La lettre choisie est dans le mot")

        hiddenRunes := []rune(hiddenWord)
        finalRunes := []rune(finalWord)

        for i := 0; i < len(finalRunes); i++ {
            if finalRunes[i] == rune(letter[0]) && hiddenRunes[i] == '_' {
                hiddenRunes[i] = rune(letter[0])
            }
        }

        updatedHiddenWord := string(hiddenRunes)
        
        return updatedHiddenWord, 0
    }
    
    return hiddenWord, 1
}
