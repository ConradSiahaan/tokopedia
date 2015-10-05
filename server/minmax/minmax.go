package minmax

/*
 * FindMinmax
 * Return Min, Max, and Sum value of a passed string parameter
 */
func FindMinmax(text string) (int, int, int) {
	
	var data = map[string]int{
	    "a": 1,
	    "b": 2,
	    "c": 3,
	    "d": 4,
	    "e": 5,
	    "f": 6,
	    "g": 7,
	    "h": 8,
	    "i": 9,
	    "j": 10,
	    "k": 11,
	    "l": 12,
	    "m": 13,
	    "n": 14,
	    "o": 15,
	    "p": 16,
	    "q": 17,
	    "r": 18,
	    "s": 19,
	    "t": 20,
	    "u": 21,
	    "v": 22,
	    "w": 23,
	    "x": 24,
	    "y": 25,
	    "z": 26,
	    "A": 1,
	    "B": 2,
	    "C": 3,
	    "D": 4,
	    "E": 5,
	    "F": 6,
	    "G": 7,
	    "H": 8,
	    "I": 9,
	    "J": 10,
	    "K": 11,
	    "L": 12,
	    "M": 13,
	    "N": 14,
	    "O": 15,
	    "P": 16,
	    "Q": 17,
	    "R": 18,
	    "S": 19,
	    "T": 20,
	    "U": 21,
	    "V": 22,
	    "W": 23,
	    "X": 24,
	    "Y": 25,
	    "Z": 26,
	}

	// var text string
	// text = "Hello World!"

	if text != "" {
		i := 1
		min := data[text[0:len(text)-(len(text) - 1)]]
		max := data[text[0:len(text)-(len(text) - 1)]]
		sum := data[text[0:len(text)-(len(text) - 1)]]

		for i < len(text) {
			if data[text[i:len(text)-((len(text) - 1)-i)]] < min && data[text[i:len(text)-((len(text) - 1)-i)]] != 0 {
				min = data[text[i:len(text)-((len(text) - 1)-i)]]
			}
			if data[text[i:len(text)-((len(text) - 1)-i)]] > max {
				max = data[text[i:len(text)-((len(text) - 1)-i)]]
			}

			sum += data[text[i:len(text)-((len(text) - 1)-i)]]

			i++
		}

		// fmt.Println("The result ", min, max, sum)
		return min, max, sum
	}else{
		return 0,0,0
	}
}

/*
 * MinMaxTemplateStr
 * Returning HTML Template For MinMax View
 */
func MinmaxTemplateStr() string {

	templateStr := `<html><head><title>Generate Min Max and Sum</title></head><body><form action="/minmax/" name=f method="GET"><input maxLength=1024 size=70 name=s value="" title="Text to Find Min, Max, & Sum" placeholder="Enter Your Strings"><input type=submit value="Generate Result" name=mm></form></body></html>`

	return templateStr
}