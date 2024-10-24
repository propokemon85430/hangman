package pendu

import (
	"fmt"
	"strings"
)

func ValidLetter() string {
	tab := CreateWord()
	vies := 10
	final_word := tab[0]
	hidden_word := tab[1]
	empty_word := ""
	fmt.Println("Choose:")
	fmt.Println(final_word)
	final_word = final_word[:len(final_word)-1]
	hidden_word = hidden_word[:len(hidden_word)-1]
	fmt.Println("You have to guess the word starting with this: ", hidden_word)
	var letter string
	for vies != 0 {
		fmt.Scanln(&letter)
		if strings.Contains(final_word, letter) {
			println("The letter chosen is in the word")
			for i := 0; i < len([]rune(final_word)); i++ {
				if string(final_word[i]) == letter {
					empty_word = empty_word + string(final_word[i])
				} else if hidden_word[i] >= 'a' && hidden_word[i] <= 'z' {
					empty_word = empty_word + string(final_word[i])
				} else {
					empty_word = empty_word + "_"
				}
			}
			if empty_word == final_word {
				fmt.Println("Congrats, you found the word which is", final_word)
				break
			}
		} else {
			println("Not present in the word, _ attempts remaining")
			vies = vies - 1
		}
		if len([]rune(empty_word)) > len([]rune(final_word)) {
			empty_word = empty_word[:len(empty_word)-1]
		}
		if len([]rune(empty_word)) != 0 {
			fmt.Println(empty_word)
			hidden_word = empty_word
			empty_word = ""
		}
		if vies == 0 {
			fmt.Println("Vous avez perdu. Le mot était :", final_word)
			return "perdu"
		}

	}
	return ("perdu")
}
