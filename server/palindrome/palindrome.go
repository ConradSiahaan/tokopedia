package palindrome

import (
	"bytes"
	"strings"
)

func FindPalindrome(text string) int {
	var data = map[string]int{
	    "a": 1,
	    "b": 1,
	    "c": 1,
	    "d": 1,
	    "e": 1,
	    "f": 1,
	    "g": 1,
	    "h": 1,
	    "i": 1,
	    "j": 1,
	    "k": 1,
	    "l": 1,
	    "m": 1,
	    "n": 1,
	    "o": 1,
	    "p": 1,
	    "q": 1,
	    "r": 1,
	    "s": 1,
	    "t": 1,
	    "u": 1,
	    "v": 1,
	    "w": 1,
	    "x": 1,
	    "y": 1,
	    "z": 1,
	    "A": 1,
	    "B": 1,
	    "C": 1,
	    "D": 1,
	    "E": 1,
	    "F": 1,
	    "G": 1,
	    "H": 1,
	    "I": 1,
	    "J": 1,
	    "K": 1,
	    "L": 1,
	    "M": 1,
	    "N": 1,
	    "O": 1,
	    "P": 1,
	    "Q": 1,
	    "R": 1,
	    "S": 1,
	    "T": 1,
	    "U": 1,
	    "V": 1,
	    "W": 1,
	    "X": 1,
	    "Y": 1,
	    "Z": 1,
	} // Var Data -> MAP

	i := 0

	var newText, newTextTemp, newA, newB string
	var buffer, bufferA, bufferB bytes.Buffer

	if text != "" {
		for i < len(text) {

			if data[text[i:len(text)-((len(text) - 1)-i)]] == 1 {
				buffer.WriteString(text[i:len(text)-((len(text) - 1)-i)])
			}
			i++

		}

		i = 0
		newTextTemp = buffer.String()
		newText = strings.ToLower(newTextTemp)
		div := len(newText) / 2
		mod := len(newText) % 2

		if mod == 1 {
			for i = len(newText) - 1; i > div; i-- {
				bufferA.WriteString(newText[i:len(newText)-((len(newText) - 1)-i)])
			}
		}else{
			for i = len(newText) - 1; i >= div; i-- {
				bufferA.WriteString(newText[i:len(newText)-((len(newText) - 1)-i)])
			}
		}

		newA = bufferA.String()

		for i = 0; i < div; i++ {
			bufferB.WriteString(newText[i:len(newText)-((len(newText) - 1)-i)])
		}

		newB = bufferB.String()

		if newA == newB {
			return 1
		} else {
			return 0
		}

	}else{
		return 0
	}
}

func PalindromeTemplateStr() string{

	templateStr := `<html><head><title>Generate Palindrome</title></head><body><form action="/palindrome/" name=f method="GET"><input maxLength=1024 size=70 name=s value="" title="Text to Find Palindrome" placeholder="Enter Your Strings"><input type=submit value="Generate Result" name=mm></form></body></html>`

	return templateStr
}