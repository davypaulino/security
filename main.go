package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
)

type VigenerePageData struct {
	Title string
	Description string
	SecretKey string
	Call string
	BtnCall string
	Btn string
}

type PageData struct {
	Title string
    Message string
	Description string
	CesarFactor int
	Call string
	BtnCall string
	CallBreak string
	Btn string
}

// HTMLs
var cesar = template.Must(template.ParseFiles("cesar.html"))
var cesarBreak = template.Must(template.ParseFiles("break-cesar.html"))
var vigenere = template.Must(template.ParseFiles("vigenere.html"))
var cryptLabel = template.Must(template.ParseFiles("crypto-label.html"))
var index = template.Must(template.ParseFiles("index.html"))

// Cripto Endpoints
func cesarEncryptHandler(w http.ResponseWriter, r *http.Request) {
	data := cesarReadRequest(r, encryptFactor)
	cryptLabel.Execute(w, data)
}

func cesarDecryptHandler(w http.ResponseWriter, r *http.Request) {
	data := cesarReadRequest(r, decryptFactor)
	cryptLabel.Execute(w, data)
}

type Response struct {
	Letters []string
	Factors []int
}

func cesarBreakHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

	message := r.URL.Query().Get("message")

    // Pegando o parâmetro "letters" e dividindo em um array
    lettersParam := r.URL.Query().Get("letters")
    letters := strings.Split(lettersParam, ",")

	runes := trackFrequency(message)
	result := generateFactors(runes, letters)

	responseData := Response{
        Letters: result.Letters,
		Factors: result.Factors,
    }

    if err := json.NewEncoder(w).Encode(responseData); err != nil {
        http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
    }
}

func vigenereEncryptHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	message := strings.TrimSpace(params.Get("message"))
	secret := strings.TrimSpace(params.Get("secret"))
	data := CipherMessage{Message: vigenereCipher(message, secret, encryptFactor)}
	cryptLabel.Execute(w, data)
}

func vigenereDecryptHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	message := strings.TrimSpace(params.Get("message"))
	secret := strings.TrimSpace(params.Get("secret"))
	data := CipherMessage{Message: vigenereCipher(message, secret, decryptFactor)}
	cryptLabel.Execute(w, data)
}

// Page Handlers
func cesarEncryptPageHandler(w http.ResponseWriter, r *http.Request) {
    data := CesarPageData{
		Title: "Criptografar Cifra de Cesar",
		Description: "",
		CesarFactor: cesarFactor,
		Call: "/cesar-lock",
		BtnCallBack: "/cesar/decrypt",
		BtnBack: "Decriptar",
		BtnCallNext: "/cesar/break",
		BtnNext: "Break",
	}
    cesar.Execute(w, data)
}

func cesarBreakPageHandler(w http.ResponseWriter, r *http.Request) {
    data := CesarPageData{
		Title: "Break Cifra de Cesar",
		Description: "",
		Call: "/cesar-break",
		BtnCallBack: "/cesar/encrypt",
		BtnBack: "Encriptar",
		BtnCallNext: "/cesar/decrypt",
		BtnNext: "Decriptar",
	}
    cesarBreak.Execute(w, data)
}

func cesarDecryptPageHandler(w http.ResponseWriter, r *http.Request) {
    data := CesarPageData{
		Title: "Desencriptação Cifra de Cesar",
		Description: "",
		CesarFactor: cesarFactor,
		Call: "/cesar-unlock",
		BtnCallBack: "/cesar/break",
		BtnBack: "Break",
		BtnCallNext: "/cesar/encrypt",
		BtnNext: "Encriptar",
	}
    cesar.Execute(w, data)
}

func vigenereEncryptPageHandler(w http.ResponseWriter, r *http.Request) {
	data := VigenerePageData{
		Title: "Criptografar Cifra de Vigenere",
		Description: "",
		SecretKey: vigenereSecret,
		Call: "/vigenere-lock",
		BtnCall: "/vigenere/decrypt",
		Btn: "Decriptar",
	}
	vigenere.Execute(w, data)
}

func vigenereDecryptPageHandler(w http.ResponseWriter, r *http.Request) {
	data := VigenerePageData{
		Title: "Criptografar Cifra de Vigenere",
		Description: "",
		SecretKey: vigenereSecret,
		Call: "/vigenere-unlock",
		BtnCall: "/vigenere/encrypt",
		Btn: "Encriptar",
	}
	vigenere.Execute(w, data)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    data := PageData{Title: "Desencriptação Cifra de Cesar"}
    index.Execute(w, data)
}

func main() {
	// Cesar
	http.HandleFunc("/cesar-lock", cesarEncryptHandler)
	http.HandleFunc("/cesar-break", cesarBreakHandler)
	http.HandleFunc("/cesar-unlock", cesarDecryptHandler)
	http.HandleFunc("/cesar/encrypt", cesarEncryptPageHandler)
	http.HandleFunc("/cesar/decrypt", cesarDecryptPageHandler)
	http.HandleFunc("/cesar/break", cesarBreakPageHandler)
	
	// Vigenere
	http.HandleFunc("/vigenere-lock", vigenereEncryptHandler)
	http.HandleFunc("/vigenere-unlock", vigenereDecryptHandler)
	http.HandleFunc("/vigenere/encrypt", vigenereEncryptPageHandler)
	http.HandleFunc("/vigenere/decrypt", vigenereDecryptPageHandler)
	
	// Index
	http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":8080", nil)
}
