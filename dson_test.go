package dson

import (
	"strings"
	"testing"
)

var (
	r string //returned
	e string //expected
)

func TestEncode(t *testing.T) {
	r = Encode(`{"foo":"bar","x":17408,"y":34  ,"you":"here","happy":true,"sad":false,"full":null,"fruits":["apple","banana","strawberry"],"quote":"\"lol\""}`)
	e = `such "foo" is "bar", "x" is 42very3, "y" is 42, "you" is "here", "happy" is yes, "sad" is no, "full" is empty, "fruits" is so "apple" and "banana" and "strawberry" many, "quote" is "\"lol\"" wow`
	check(t)
}

func TestDecode(t *testing.T) {
	r = Decode(`such "foo" is "bar", "x" is 42very3, "y" is 42 , "you" is "here", "happy" is yes, "sad" is no, "full" is empty, "fruits" is so "apple" and "banana" and "strawberry" many, "quote" is "\"lol\"" wow`)
	e = `{"foo":"bar","x":17408,"y":34,"you":"here","happy":true,"sad":false,"full":null,"fruits":["apple","banana","strawberry"],"quote":"\"lol\""}`
	check(t)
}

func TestValid(t *testing.T) {
	if !Valid(`such "foo" is "bar" wow`) {
		t.Fatal("Unable to validate DSON")
	}
	if Valid(`such "foo" is "bar"`) {
		t.Fatal("Unable to detect faulty DSON")
	}
}

func TestMarshal(t *testing.T) {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	r, err := Marshal(group)
	if err != nil || !Valid(r) {
		msg := "Unable to marshal DSON"
		if !Valid(r) {
			msg += ", invalid dson -- " + Decode(r)
		}
		t.Fatal(msg, "\n", r)
	}
}

func TestUnmarshal(t *testing.T) {
	d := `so sUch "Name" is "Platypus" aNd "Order" IS "Monotremata" wow And sUch "Name" is "Quoll" AND "Order" is "Dasyuromorphia" wow mAny`
	if !Valid(d) {
		t.Fatal("DSON is not valid")
	}
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := Unmarshal(d, &animals)
	if err != nil {
		t.Fatalf("Unable to unmarshal DSON: %+v", animals)
	}
}

func check(t *testing.T) {
	r = strings.Replace(r, "!", ",", -1)
	r = strings.Replace(r, ".", ",", -1)
	r = strings.Replace(r, "?", ",", -1)
	r = strings.Replace(r, "also", "and", -1)
	if r != e {
		t.Fatalf(`Expected "%s" but got "%s" `, e, r)
	}
}
