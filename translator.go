package dson

var (
	quote    bool
	array    bool
	escaping bool
	token    string
	current  string
	output   string
	number   string
	i        int
)

func reset() {
	quote = false
	array = false
	escaping = false
	token = ""
	current = ""
	output = ""
	number = ""
}

func addSeparator(currentChar string) string {
	if array {
		currentChar = getRandom(" and", " also")
	} else {
		currentChar = getRandom(".", ",", "!", "?")
	}
	return currentChar + " "
}

func addQuote(currentChar string, isDSON bool) {
	output += "\"" + current + "\""
	if !isDSON {
		output += " "
	}
	quote = false
	current = ""
}

func checkForArray(currentChar string) {
	if trimSpace(currentChar) == "" {
		return
	}
	token += currentChar
	if currentChar == "]" {
		array = false
	}
	if currentChar == "[" {
		array = true
	}
}

func translate(data []rune, keywords map[string]string, isDSON bool) string {
	reset()
	for i = 0; i < len(data); i++ {
		currentChar := string(data[i])
		if currentChar == "," && !isDSON {
			output = trimSpace(output) + addSeparator(currentChar)
			continue
		}
		if quote {
			if (currentChar == "\"" || currentChar == "'") && !escaping {
				addQuote(currentChar, isDSON)
				continue
			}
			escaping = currentChar == "\\"
			current += currentChar
			continue
		}
		if currentChar == "\"" {
			quote = true
			current = ""
			continue
		}
		checkForArray(currentChar)
		if isNum(token) || contains(token, "very") {
			number += token
			token = ""
			for string(data[i+1]) == " " {
				i++
			}
			if isSeparator(data[i+1]) {
				if !isDSON {
					output += generateVery(number)
				} else {
					output += getDecimal(parseNumber(number))
				}
				number = ""
			}
			continue
		}
		if lToken := toLower(token); exists(keywords[lToken]) {
			output += keywords[lToken]
			if !isDSON {
				output += " "
			}
			token = ""
		}
	}
	return trimSpace(output)
}
