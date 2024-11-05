package pendu

import (
    "fmt"
    "strings"
)

// ValidLetter vérifie si la lettre proposée est dans le mot final
func ValidLetter(finalWord, hiddenWord, letter string) (string, int) {
    // Vérification si la lettre est présente dans le mot final
    if strings.Contains(finalWord, letter) {
        fmt.Println("La lettre choisie est dans le mot")
        // Remplacer le premier point (".") par la lettre dans le mot caché
        for i := 0; i < len(finalWord); i++ {
            if string(finalWord[i]) == letter {
                hiddenWord = strings.Replace(hiddenWord, ".", string(finalWord[i]), 1)
            }
        }
        return hiddenWord, 0 // Retourner le mot caché mis à jour
    }

    fmt.Println("La lettre n'est pas présente dans le mot")
    return hiddenWord, 1 // Retourner le mot caché sans modification
}
