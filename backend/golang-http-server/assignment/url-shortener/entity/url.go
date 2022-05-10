package entity

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Buatlah entity URL dengan atribut LongURL dan ShortURL masing-masing string

// TODO: answer here

type URL struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

func GetRandomShortURL(longURL string) string {
	showL := fmt.Sprintf("%s %d", longURL, time.Now().Unix())
	sum := sha256.Sum256([]byte(showL))
	hashString := fmt.Sprintf("%x", sum)
	return hashString[:8]
}
