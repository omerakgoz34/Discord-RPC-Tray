package main

import (
	"log"
)

// Lang - Language Strings
var Lang = map[string]string{}

// InitLang - Sets the app language
func InitLang() {
	switch Config["lang"] {
	case "en":
		Lang = LangEN
	case "tr":
		Lang = LangTR
	default:
		Lang = LangEN
	}
	log.Println("Lang: ", Lang)
}
