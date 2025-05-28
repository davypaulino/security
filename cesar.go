package main

import (
	"encoding/json"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CesarPageData struct {
	Title string
	Description string
	CesarFactor int
	Call string
	BtnCallNext string
	BtnCallBack string
	BtnNext string
	BtnBack string
}

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

type CipherRequest struct {
    Message string `json:"message"`
    Factor  string `json:"factor"`
}

func cesarReadRequest(r *http.Request, action func(rune, rune, rune) string) CipherMessage {
    var req CipherRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        return CipherMessage{}
    }

    message := strings.TrimSpace(req.Message)
    if message == "" {
        return CipherMessage{}
    }

    newFactor, err := strconv.Atoi(req.Factor)
    if err != nil {
        newFactor = cesarFactor
    }

	cesarMessage := cesarCipher(message, action, abs(newFactor))
    data := CipherMessage{Message: cesarMessage}
    return data
}

func trackFrequency(message string) map[rune] int {
	runes := make(map[rune]int)
	for _, c := range strings.ToUpper(message) {
        if c >= 'A' && c <= 'Z' {
            runes[c]++
        }
    }
	return runes
}

type FactorsCesar struct {
    Factors []int
    Letters []string
}

func generateFactors(runes map[rune]int, letters []string) FactorsCesar {
    type pair struct {
        letter rune
        freq  int
    }
    pairs := make([]pair, 0, len(runes))

    for letter, freq := range runes {
        pairs = append(pairs, pair{letter, freq})
    }

    sort.SliceStable(pairs, func(i, j int) bool {
        return pairs[i].freq > pairs[j].freq
    })

    moreFreqLetters := make([]string, 0, len(pairs))
    for _, p := range pairs {
        moreFreqLetters = append(moreFreqLetters, string(p.letter))
    }

    result := make([]int, len(letters))
    for i := range letters {
        if len(letters[i]) > 0 {
            result[i] = int(moreFreqLetters[i][0]) - int(letters[i][0])
        }
    }

    return FactorsCesar{Factors: result, Letters: moreFreqLetters[:len(letters)]}
}