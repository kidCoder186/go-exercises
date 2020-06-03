package cipher

import (
	"strings"
	"unicode"
)

// ShiftCipher structure.
type ShiftCipher struct {
	distance int
}

// VigenereCipher structure.
type VigenereCipher struct {
	key string
}

// Encode returns a cipher text of plain text.
func (sc ShiftCipher) Encode(plain string) string {
	var cipher []rune
	runes := []rune(strings.ToLower(plain))
	for _, r := range runes {
		if unicode.IsLetter(r) {
			c := rune('a') + (r-rune('a')+rune(sc.distance))%26
			cipher = append(cipher, c)
		}
	}
	return string(cipher)
}

// Decode returns a plain text of cipher text.
func (sc ShiftCipher) Decode(cipher string) string {
	return NewShift(-sc.distance).Encode(cipher)
}

// Encode returns a cipher text of plain text.
func (vc VigenereCipher) Encode(plain string) string {
	plain = strings.Join(strings.FieldsFunc(plain, func(r rune) bool {
		return !unicode.IsLetter(r)
	}), "")
	plain = strings.ToLower(plain)
	var temp string
	for len(plain) > len(temp) {
		temp += vc.key
	}
	keyCipher := string(temp[:len(plain)])
	var cipher []byte
	for i := range plain {
		dist := (keyCipher[i] - 'a' + 26) % 26
		c := 'a' + (plain[i]-'a'+dist)%26
		cipher = append(cipher, c)
	}
	return string(cipher)
}

// Decode returns a plain text of cipher text.
func (vc VigenereCipher) Decode(cipher string) string {
	newKey := []byte(vc.key)
	for i := range vc.key {
		newKey[i] = 'a' + ('a'+26-vc.key[i])%26
	}
	return NewVigenere(string(newKey)).Encode(cipher)
}

// NewCaesar returns a ShiftCipher object with distance = 3.
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift returns a ShiftCiper object.
// If pass a invalid distance then `nil` will be returned
func NewShift(distance int) Cipher {
	if distance <= -26 || distance >= 26 || distance == 0 {
		return nil
	}
	distance = (distance + 26) % 26
	return ShiftCipher{distance}
}

// NewVigenere returns a VigenereCipher object.
// If pass a invalid key then `nil` will be returned
func NewVigenere(key string) Cipher {
	var valid bool
	for _, c := range key {
		if c > 'a' {
			valid = true
		}
		if !unicode.IsLower(rune(c)) {
			return nil
		}
	}
	if !valid {
		return nil
	}
	return VigenereCipher{key}
}
