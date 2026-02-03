package main

import (
	"fmt"
	"time"
)

func getRandomTable() map[rune]rune {
	subTable := map[rune]rune{}
	allowedLetters := map[rune]bool{}

	for i := 'n'; i <= 'z'; i++ {
		allowedLetters[i] = true
	}

	for i := 'a'; i <= 'm'; i++ {
		for {
			var val rune = rune('n' + (time.Now().UnixMilli() % 13))
			if allowedLetters[val] == true {
				subTable[i] = val
				allowedLetters[val] = false
				break
			}
		}

	}
	return subTable
}

func macRandom(subTable map[rune]rune, text string) string {
	result := ""
	for _, letter := range text {
		value, found := subTable[letter]
		if found {
			result += string(value)
		} else {
			for k, v := range subTable {
				if v == letter {
					result += string(k)

				}
			}
		}

	}

	return result
}

func task4() {
	var input string
	fmt.Println("Enter Text to be encrypted:")
	fmt.Scanf("%s", &input)
	subTable := getRandomTable()
	encrypted := macRandom(subTable, input)
	decrypted := macRandom(subTable, encrypted)
	fmt.Println("Encrypted: " + encrypted)
	fmt.Println("Decrypted: " + decrypted)
}
