package strings

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"unicode"
)

// Capitalize make first letter to upper
// Source https://www.cnblogs.com/Detector/p/9686443.html
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// IsCapitalize whether the first letter is upper
func IsCapitalize(s string) bool {
	if len(s) < 1 {
		return false
	}
	return unicode.IsUpper([]rune(s)[0])
}

// SplitToChunks Split large string in n-size chunks
func SplitToChunks(s string, chunkSize int) []string {
	if chunkSize < 1 {
		return []string{}
	}

	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}

// Rand generate rand string with bytes size
func Rand(size int) (string, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// IsNumber determine if it is a pure numeric string
func IsNumber(str string) bool {
	if len(str) < 1 {
		return false
	}
	for _, r := range str {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

// RemoveEmpties move all empties
func RemoveEmpties(s string) string {
	return Remove(s, " ")
}

// Remove remove all old from s
func Remove(s, old string) string {
	return strings.Replace(s, old, "", -1)
}
