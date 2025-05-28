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
	secretSize := len(secret)
	m = strings.ToUpper(m)
	secret = strings.ToUpper(secret)

    for i := range messageSize {
		c := rune(m[i])
        
		secretChar := rune(secret[i % secretSize])

        if c >= 'A' && c <= 'Z' && unicode.IsLetter(secretChar) {
            result += action('A', c, (secretChar - 'A') % 26)
		} else {
            result += string(c)
        }
    }

    return result
}