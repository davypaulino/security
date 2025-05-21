package main

import (
    "html/template"
    "net/http"
	"os"
	"strconv"
)

type PageData struct {
	Title string
    Message string
	Call string
	BtnCall string
	Btn string
}

// HTMLs
var cesar = template.Must(template.ParseFiles("cesar.html"))
var cesarCipher = template.Must(template.ParseFiles("cesar-cipher.html"))
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
		} else if c >= 'a' && c <= 'z' {
            encrypted += string(( (c - 'a' + rune(cesarFactor)) % 26 + 'a'))
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
            decrypted += string(( (c - 'A' - rune(cesarFactor) + 26) % 26 + 'A'))
		} else if c >= 'a' && c <= 'z' {
            decrypted += string(( (c - 'a' - rune(cesarFactor) + 26) % 26 + 'a'))
        } else {
            decrypted += string(c)
        }
    }

    return decrypted
}

// Handler Html
func cesarEncryptHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	cesarMessage := encryptCesar(params.Get("message"))
	data := PageData{Message: cesarMessage}
	cesarCipher.Execute(w, data)
}

func cesarDecryptHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	cesarMessage := decryptCesar(params.Get("message"))
	data := PageData{Message: cesarMessage}
	cesarCipher.Execute(w, data)
}

func cesarEncryptPageHandler(w http.ResponseWriter, r *http.Request) {
    data := PageData{
		Title: "Criptografar Cifra de Cesar",
		Call: "/cesar-lock",
		BtnCall: "/cesar/decrypt",
		Btn: "Decriptar",
	}
    cesar.Execute(w, data)
}

func cesarDecryptPageHandler(w http.ResponseWriter, r *http.Request) {
    data := PageData{
		Title: "Desencriptação Cifra de Cesar",
		Call: "/cesar-unlock",
		BtnCall: "/cesar/encrypt",
		Btn: "Encriptar",
	}
    cesar.Execute(w, data)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    data := PageData{Title: "Desencriptação Cifra de Cesar"}
    index.Execute(w, data)
}

func main() {
	http.HandleFunc("/cesar-lock", cesarEncryptHandler)
	http.HandleFunc("/cesar-unlock", cesarDecryptHandler)
	http.HandleFunc("/cesar/encrypt", cesarEncryptPageHandler)
	http.HandleFunc("/cesar/decrypt", cesarDecryptPageHandler)
	http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":8080", nil)
}
