// Package dson implements encoding, decoding, marshaling, unmarshaling,
// and verification of the DSON (Doge Serialized Object Notation)
// as defined here: https://dogeon.xyz/.
//
// Examples can be found here: https://muzzammil.xyz/dson.go
//
// Version 1.0.0
//
// MIT License;
// https://github.com/muhammadmuzzammil1998/dsongo/blob/master/LICENSE
//
// (c) 2018 Muhammad Muzzammil <email@muzzammil.xyz>
package dson

import "encoding/json"

// Encode encodes a JSON string into the DSON equivalent and
// retruns that DSON string
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

// Decode decodes a DSON string into the JSON equivalent and
// retruns that JSON string
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

// Marshal returns the DSON encoding of v.
func Marshal(v interface{}) (string, error) {
	m, err := json.Marshal(v)
	return Encode(string(m)), err
}

// Unmarshal parses the DSON-encoded data and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// Unmarshal returns an InvalidUnmarshalError.
func Unmarshal(dson string, v interface{}) error { return json.Unmarshal([]byte(Decode(dson)), &v) }

// Valid returns true if the input DSON string is valid, otherwise returns false
func Valid(dson string) bool { return json.Valid([]byte(Decode(dson))) }
