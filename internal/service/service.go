package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertMorseOrText(s string) string {
	if isMorseCode(s) {
		return morse.ToText(s)
	}

	return morse.ToMorse(s)
}

func isMorseCode(s string) bool {
	if strings.TrimSpace(s) == "" {
		return false
	}

	for _, char := range s {
		if char != '-' && char != ' ' && char != '.' {
			return false
		}
	}

	return true
}
