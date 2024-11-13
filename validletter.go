package pendu

import (
    "fmt"
    "strings"
)

func ValidLetter(finalWord, hiddenWord, letter string) (string, int) {
    // Assurer que hiddenWord a la même longueur que finalWord
    if len(hiddenWord) > len(finalWord) {
        hiddenWord = hiddenWord[:len(finalWord)] // Tronquer hiddenWord si nécessaire
    }

    // Afficher les longueurs pour déboguer
    fmt.Println("Longueur de hiddenWord : ", len(hiddenWord))
    fmt.Println("Longueur de finalWord : ", len(finalWord))

    // Vérification si la lettre est dans le mot final
    if strings.Contains(finalWord, letter) {
        fmt.Println("La lettre choisie est dans le mot")

        // Convertir hiddenWord et finalWord en slices de runes pour pouvoir modifier hiddenWord
        hiddenRunes := []rune(hiddenWord)
        finalRunes := []rune(finalWord)

        // Modifier hiddenWord pour révéler la lettre correcte
        for i := 0; i < len(finalRunes); i++ {
            if finalRunes[i] == rune(letter[0]) && hiddenRunes[i] == '_' {
                hiddenRunes[i] = rune(letter[0])
            }
        }

        // Retourner hiddenWord mis à jour
        updatedHiddenWord := string(hiddenRunes)

        // Afficher la version mise à jour de hiddenWord pour le débogage
        fmt.Println("hiddenWord mis à jour : ", updatedHiddenWord)

        return updatedHiddenWord, 0
    }

    fmt.Println("La lettre n'est pas présente dans le mot")
    return hiddenWord, 1
}
