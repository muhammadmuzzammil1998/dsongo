package dson

import "encoding/json"

func Encode(json string) string {
	return translate([]rune(json), map[string]string{
		"{":     "such",
		"}":     "wow",
		":":     "is",
		"[":     "so",
		"]":     "many",
		"true":  "yes",
		"false": "no",
		"null":  "empty",
	}, false)
}

func Decode(dson string) string {
	return translate([]rune(dson), map[string]string{
		"such":  "{",
		"wow":   "}",
		"is":    ":",
		"so":    "[",
		"many":  "]",
		"and":   ",",
		"also":  ",",
		".":     ",",
		",":     ",",
		"!":     ",",
		"?":     ",",
		"yes":   "true",
		"no":    "false",
		"empty": "null",
	}, true)
}

func Marshal(v interface{}) (string, error) {
	m, err := json.Marshal(v)
	return Encode(string(m)), err
}

func Unmarshal(dson string, v interface{}) error { return json.Unmarshal([]byte(Decode(dson)), &v) }

func Valid(dson string) bool { return json.Valid([]byte(Decode(dson))) }
