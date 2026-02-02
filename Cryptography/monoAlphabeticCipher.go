package main

import "fmt"

func mac(text string) string {
	subTable := map[rune]rune{
		'a': 'q',
		'b': 'w',
		'c': 'x',
		'd': 'r',
		'e': 't',
		'f': 'y',
		'g': 'u',
		'h': 's',
		'i': 'o',
		'j': 'p',
		'k': 'n',
		'l': 'v',
		'm': 'z',
	}
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

func task3() {
	var input string
	fmt.Scanf("%s", &input)
	encrypted := mac(input)
	decrypted := mac(encrypted)
	print("Encrypted: " + encrypted + "\n")
	print("Decrypted: " + decrypted + "\n")
}
