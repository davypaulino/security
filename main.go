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
	BtnCallBack string
	BtnBack string
	BtnCallNext string
	BtnNext string
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

func cesarEncryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	data := cesarReadRequest(r, encryptFactor)
	if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
    }
}

func cesarDecryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	data := cesarReadRequest(r, decryptFactor)
	if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
    }
}

type Request struct {
    Message string   `json:"message"`
    Letters []string `json:"letters"`
}

type Response struct {
    Letters []string `json:"letters"`
    Factors []int    `json:"factors"`
}

func cesarBreakHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
        return
    }

    var req Request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao ler JSON da requisição", http.StatusBadRequest)
        return
    }
	
    if req.Message == "" || len(req.Letters) == 0 {
		http.Error(w, "Campos 'message' e 'letters' são obrigatórios", http.StatusBadRequest)
        return
    }
	
	w.Header().Set("Content-Type", "application/json")
	runes := trackFrequency(req.Message)
	result := generateFactors(runes, req.Letters)

	responseData := Response{
        Letters: result.Letters,
		Factors: result.Factors,
    }

    if err := json.NewEncoder(w).Encode(responseData); err != nil {
        http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
    }
}

func vigenereEncryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
        return
    }

    var req VigenereRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao ler JSON da requisição", http.StatusBadRequest)
        return
    }

	message := strings.TrimSpace(req.Message)
	secret := strings.TrimSpace(req.Secret)

	data := CipherMessage{Message: vigenereCipher(message, secret, encryptFactor)}
    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
    }
}

func vigenereDecryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
        return
    }

    var req VigenereRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao ler JSON da requisição", http.StatusBadRequest)
        return
    }

	message := strings.TrimSpace(req.Message)
	secret := strings.TrimSpace(req.Secret)

	data := CipherMessage{Message: vigenereCipher(message, secret, decryptFactor)}
    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
    }
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
		BtnCallBack: "/vigenere/decrypt",
		BtnBack: "Decriptar",
		BtnCallNext: "/vigenere/decrypt",
		BtnNext: "Decriptar",
	}
	vigenere.Execute(w, data)
}

func vigenereDecryptPageHandler(w http.ResponseWriter, r *http.Request) {
	data := VigenerePageData{
		Title: "Criptografar Cifra de Vigenere",
		Description: "",
		SecretKey: vigenereSecret,
		Call: "/vigenere-unlock",
		BtnCallBack: "/vigenere/encrypt",
		BtnBack: "Encriptar",
		BtnCallNext: "/vigenere/encrypt",
		BtnNext: "Encriptar",
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
