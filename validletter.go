package pendu

import (
    "fmt"
    "strings"
)

func ValidLetter(finalWord, hiddenWord, letter string) (string, int) {
    if letter>='A' || letter <='Z' {
        letter=strings.ToLower(letter) }
    if strings.Contains(finalWord, letter) {
        fmt.Println("La lettre choisie est dans le mot")

        hiddenWordArr := []rune(hiddenWord)

        for i := 0; i < len(finalWord); i++ {
            if rune(finalWord[i]) == rune(letter[0]) && hiddenWordArr[i] == '.' {
                hiddenWordArr[i] = rune(letter[0]) 
            }
        }

        return string(hiddenWordArr), 0 
    }

    fmt.Println("La lettre n'est pas présente dans le mot")
    return hiddenWord, 1 
}
