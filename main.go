package main

import (
    "html/template"
    "net/http"
	"os"
	"strconv"
	"strings"
)

type PageData struct {
	Title string
    Message string
}

// HTMLs
var cesarEncrypt = template.Must(template.ParseFiles("encrypt-cesar.html"))
var cesarCypher = template.Must(template.ParseFiles("cesar-label.html"))
var index = template.Must(template.ParseFiles("index.html"))

// Pegando os Valores das váriaveis de ambiente
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

// encrypt e decrypt da cifra de cesar
func encryptCesar(m string) string {
    encrypted := ""
    for _, c := range m {
        if c >= 'A' && c <= 'Z' {
            encrypted += string(( (c - 'A' + rune(cesarFactor)) % 26 + 'A'))
        } else {
            encrypted += string(c)
        }
    }

    return encrypted
}

func decryptCesar(m string) string {
    decrypted := ""
    for _, c := range m {
        if c >= 'A' && c <= 'Z' {
            decrypted += string(( (c - 'A' - rune(cesarFactor)) % 26 + 'A'))
        } else {
            decrypted += string(c)
        }
    }

    return decrypted
}

// Handler Html
func cesarCypherHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	cesarMessage := encryptCesar(strings.ToUpper(params.Get("message")))
	data := PageData{Message: cesarMessage}
	cesarCypher.Execute(w, data)
}

func cesarEncryptHandler(w http.ResponseWriter, r *http.Request) {
    data := PageData{Title: "Criptografar Cifra de Cesar"}
    cesarEncrypt.Execute(w, data)
}

func cesarDecryptHandler(w http.ResponseWriter, r *http.Request) {
    data := PageData{Title: "Desencriptação Cifra de Cesar"}
    cesarEncrypt.Execute(w, data)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    data := PageData{Title: "Desencriptação Cifra de Cesar"}
    index.Execute(w, data)
}

func main() {
    //http.HandleFunc("/", handler)
	http.HandleFunc("/cesar-label", cesarCypherHandler)
	http.HandleFunc("/cesar/encrypt", cesarEncryptHandler)
	http.HandleFunc("/cesar/decrypt", cesarDecryptHandler)
	http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/cesar/decrypt", cesarCypherHandler)
    http.ListenAndServe(":8080", nil)
}
