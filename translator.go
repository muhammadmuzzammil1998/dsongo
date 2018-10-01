package dson

func translate(data []rune, keywords map[string]string, isDSON bool) string {
	var (
		quote    bool
		array    bool
		escaping bool
		token    string
		current  string
		output   string
		number   string
	)
	for i := 0; i < len(data); i++ {
		currentChar := string(data[i])
		if currentChar == "," && !isDSON {
			if array {
				currentChar = getRandom(" and", " also")
			} else {
				currentChar = getRandom(".", ",", "!", "?")
			}
			output = trimSpace(output) + currentChar + " "
			continue
		}
		if quote {
			if (currentChar == "\"" || currentChar == "'") && !escaping {
				output += "\"" + current + "\""
				if !isDSON {
					output += " "
				}
				quote = false
				current = ""
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
		if trimSpace(currentChar) != "" {
			token += currentChar
			if currentChar == "]" {
				array = false
			}
			if currentChar == "[" {
				array = true
			}
		}
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
