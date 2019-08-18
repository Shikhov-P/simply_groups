package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unicode"
)

func ParseError(err error) {
	if err != nil {
		fmt.Println(err)
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

func IsWord(word string) bool {
	for _, letter := range word {
		if !unicode.IsLetter(letter){
			return false
		}
	}
	return true
}
