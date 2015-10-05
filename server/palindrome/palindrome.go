package palindrome

import (
	"strings"
	"log"
	"regexp"
)

func FindPalindrome(text string) int {

	if text != "" {

		reg, err := regexp.Compile("[^A-Za-z]+")
		if err != nil {
			log.Fatal(err)
		}

		newText := reg.ReplaceAllString(text,"")
		newText = strings.ToLower(strings.Trim(newText, "-"))

		var newA, newB string

		div := len(newText) / 2
		mod := len(newText) % 2
		length := len(newText)

		if mod == 1 {
			newA = newText[div+1:length]
		}else{
			newA = newText[div:length]
		}

		newB = newText[0:div]
		newARev := Reverse(newA)

		if newARev == newB {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
	
}

func PalindromeTemplateStr() string{

	templateStr := `<html><head><title>Generate Palindrome</title></head><body><form action="/palindrome/" name=f method="GET"><input maxLength=1024 size=70 name=s value="" title="Text to Find Palindrome" placeholder="Enter Your Strings"><input type=submit value="Generate Result" name=mm></form></body></html>`

	return templateStr
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}