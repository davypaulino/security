package main

import (
    "testing"
)

// Teste aprimorado para a função encryptFactor
func TestEncryptFactor(t *testing.T) {
    tests := []struct {
        start  rune
        c      rune
        factor rune
        want   string
    }{
        {'a', 'a', 3, "d"},
        {'a', 'z', 1, "a"},
        {'A', 'B', 2, "D"},
        {'A', 'Z', 3, "C"},
    }

    for _, tt := range tests {
        got := encryptFactor(tt.start, tt.c, tt.factor)
        if got != tt.want {
            t.Errorf("encryptFactor(%q, %q, %q) = %q; esperado %q", tt.start, tt.c, tt.factor, got, tt.want)
        }
    }
}


// Teste para a função decryptFactor
func TestDecryptFactor(t *testing.T) {
    tests := []struct {
        start  rune
        c      rune
        factor rune
        want   string
    }{
        {'a', 'd', 3, "a"},
        {'a', 'a', 1, "z"},
        {'A', 'D', 2, "B"},
        {'A', 'C', 3, "Z"},
    }

    for _, tt := range tests {
        got := decryptFactor(tt.start, tt.c, tt.factor)
        if got != tt.want {
            t.Errorf("decryptFactor(%q, %q, %q) = %q; esperado %q", tt.start, tt.c, tt.factor, got, tt.want)
        }
    }
}

// Teste aprimorado para cesarCipher
func TestCesarCipher(t *testing.T) {
    tests := []struct {
        m      string
        action func(rune, rune, rune) string
        factor int
        want   string
    }{
        {"abc", encryptFactor, 3, "def"},
        {"xyz", encryptFactor, 2, "zab"},
        {"DEF", decryptFactor, 3, "ABC"},
        {"MNO", encryptFactor, 5, "RST"},
        {"HELLO", decryptFactor, 7, "AXEEH"},
        {"", encryptFactor, 3, ""},
        {"123!@#", encryptFactor, 5, "123!@#"},
        {"abcdefghijklmnopqrstuvwxyz", encryptFactor, 25, "zabcdefghijklmnopqrstuvwxy"},
        {"Looooooooooooooong text", encryptFactor, 4, "Psssssssssssssssrk xibx"},
    }

    for _, tt := range tests {
        got := cesarCipher(tt.m, tt.action, tt.factor)
        if got != tt.want {
            t.Errorf("cesarCipher(%q, %d) = %q; esperado %q", tt.m, tt.factor, got, tt.want)
        }
    }
}

// Teste aprimorado para vigenereCipher
func TestVigenereCipher(t *testing.T) {
    tests := []struct {
        m      string
        action func(rune, rune, rune) string
        key    string
        want   string
    }{
        {"HELLO", encryptFactor, "KEY", "RIJVS"},
        {"RIJVS", decryptFactor, "KEY", "HELLO"},
        {"ATTACKATDAWN", encryptFactor, "LEMON", "LXFOPVEFRNHR"},
        {"LXFOPVEFRNHR", decryptFactor, "LEMON", "ATTACKATDAWN"},
        {"SHORT", encryptFactor, "LONGKEYWORD", "DVBXD"}, // Mensagem menor que a chave
        {"LONGMESSAGE", encryptFactor, "KEY", "VSLQQCCWYQI"}, // Chave menor que a mensagem
        {"VALID", encryptFactor, "K3Y!", "FAJIN"}, // Chave com caracteres inválidos (mantém mensagem inalterada)
        {"TEST", encryptFactor, "", ""}, // Chave vazia deve manter mensagem inalterada
        {"MATCH", encryptFactor, "MATCH", "YAMEO"}, // Chave igual à mensagem
        {"hello", encryptFactor, "key", "RIJVS"}, // Teste com letras minúsculas
        {"HELLOworld", encryptFactor, "KEY", "RIJVSUYVJN"}, // Teste com mistura de maiúsculas e minúsculas
        {"MixedCASE", encryptFactor, "Secret", "EMZVHVSWG"}, // Outro caso misto
        {"MixedCASE", encryptFactor, "1", "MIXEDCASE"}, // Outro caso misto
    }

    for _, tt := range tests {
        got := vigenereCipher(tt.m, tt.key, tt.action)
        if got != tt.want {
            t.Errorf("vigenereCipher(%q, %q) = %q; esperado %q", tt.m, tt.key, got, tt.want)
        }
    }
}
