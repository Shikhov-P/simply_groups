package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"unicode"
)

func ParseError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	ParseError(err)
}

func IsAlphaNumeric(str string) bool {
	words := strings.Split(str, " ")
	for _, word := range words {
		for _, symbol := range word {
			if !(unicode.IsLetter(symbol) || unicode.IsNumber(symbol)) {
				return false
			}
		}
	}
	return true
}
