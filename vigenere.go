package main

import (
	"os"
	"strings"
	"unicode"
)

type VigenereRequest struct {
    Message string `json:"message"`
    Secret  string `json:"secret"`
}

var vigenereSecret = func() string {
    envSecret := os.Getenv("VIGENERE_SECRET")
    secret := "a"

    if envSecret != "" {
        secret = envSecret
    }

    return secret
}()

func vigenereCipher(m, secret string, action func(rune, rune, rune) string) string {
	result := ""
	if (strings.TrimSpace(secret) == "" || strings.TrimSpace(m) == "") {
		return result;
	}
	
	messageSize := len(m)
	m = strings.ToUpper(m)
	secret = strings.ToUpper(secret)

	for messageSize > len(secret) {
		secret += secret
	}
	secret = secret[:messageSize]

    for i := range messageSize {
		c := rune(m[i])
        
		secretChar := rune(secret[i])

        if c >= 'A' && c <= 'Z' && unicode.IsLetter(secretChar) {
            result += action('A', c, (secretChar - 'A') % 26)
		} else {
            result += string(c)
        }
    }

    return result
}