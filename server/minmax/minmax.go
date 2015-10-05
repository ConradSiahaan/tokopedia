package minmax

import (
	"strings"
	"log"
	"regexp"
)

/*
 * FindMinmax
 * Return Min, Max, and Sum value of a passed string parameter
 */
func FindMinmax(text string) (int, int, int) {
	
	if text == "" {
		return 0,0,0
	}
	
	reg, err := regexp.Compile("[^A-Za-z]+")
	if err != nil {
		log.Fatal(err)
	}

	newText := reg.ReplaceAllString(text,"")
	newText = strings.ToLower(strings.Trim(newText, "-"))

	i := 1
	temp := 0

	r := []rune(newText)

	min := int(r[0]) - 96
	max := int(r[0]) - 96
	sum := int(r[0]) - 96

	for i < len(newText) {
		temp = int(r[i]) - 96
		if temp < min {
			min = temp
		}
		if temp > max {
			max = temp
		}

		sum += temp

		i++

	}

	return min, max, sum
}

/*
 * MinMaxTemplateStr
 * Returning HTML Template For MinMax View
 */
func MinmaxTemplateStr() string {

	templateStr := `<html><head><title>Generate Min Max and Sum</title></head><body><form action="/minmax/" name=f method="GET"><input maxLength=1024 size=70 name=s value="" title="Text to Find Min, Max, & Sum" placeholder="Enter Your Strings"><input type=submit value="Generate Result" name=mm></form></body></html>`

	return templateStr
}