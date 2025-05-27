package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"
)

var cesarFactor = func() int {
    factor := os.Getenv("CESARCYPHER")
    shift := 3

    if factor != "" {
        parsedFactor, err := strconv.Atoi(factor)
        if err == nil {
            shift = parsedFactor
        }
    }

    return shift
}()

func cesarCipher(m string, action func(rune, rune, rune) string, factor int) string {
    result := ""
    for _, c := range m {
        if c >= 'A' && c <= 'Z' {
            result += action('A', c, rune(factor))
		} else if c >= 'a' && c <= 'z' {
			result += action('a', c, rune(factor))
        } else {
            result += string(c)
        }
    }

    return result
}

func cesarReadRequest(r *http.Request, action func(rune, rune, rune) string) PageData {
    params := r.URL.Query()
	message := strings.TrimSpace(params.Get("message"))
	if message == "" {
		return PageData{}
	}

	newFactor, err := strconv.Atoi(params.Get("factor"))
	if (err != nil) {
		newFactor = cesarFactor
	}

	cesarMessage := cesarCipher(message, action, abs(newFactor))
    data := PageData{Message: cesarMessage}
    return data
}