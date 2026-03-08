package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"unicode"
)

func GenerateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func Slugify(s string) string {
	var sb strings.Builder
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(r)
		} else if unicode.IsSpace(r) || r == '-' {
			sb.WriteRune('-')
		}
	}
	return strings.Trim(sb.String(), "-")
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func Ptr[T any](v T) *T {
	return &v
}
