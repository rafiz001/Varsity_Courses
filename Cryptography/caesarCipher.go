package main

import "fmt"

func cs(text string, key int) string {
	result := ""
	for _, ch := range text {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			is_upper := (ch >= 'A' && ch <= 'Z')
			var b rune
			if is_upper {
				b = 'A'
			} else {
				b = 'a'
			}
			var keyed_char = (ch-b+rune(key)+26)%26 + b
			result += string(keyed_char)
		} else {
			result += string(ch)
		}
	}
	return result
}
func csd(text string, key int) string {
	return cs(text, -key)
}

func task1() {
	var input rune
	fmt.Scanf("%c", &input)
	fmt.Println(string(input))
	fmt.Println(string(input - ('a' - 'A')))
}
func task2() {
	var input string
	var key int
	fmt.Scanf("%s", &input)
	fmt.Scanf("%d", &key)
	encrypted := cs(input, key)
	decrypted := csd(encrypted, key)
	print("Encrypted: " + encrypted + "\n")
	println("Decrypted: " + decrypted + "\n")
}
