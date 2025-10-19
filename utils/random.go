package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

func IsEmpty(data string) bool {
	if len(data) <= 0 {
		return true
	} else {
		return false
	}

}

func LeftRatation(input string, step int) string {
	if IsEmpty(input) {
		return ""
	}
	res := ""
	if step > len(input) {
		step = step % len(input)
	}
	res += input[len(input)-step:]
	res += input[:len(input)-step]
	return res
}

func GenerateSalt() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", errors.New("read secure random data")
	}
	return hex.EncodeToString(bytes), nil
}
