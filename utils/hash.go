package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

const (
	// Pbkdf2Iterations sets the amount of iterations used by the PBKDF2 hashing algorithm
	Pbkdf2Iterations int = 15000
	// HashBytes sets the amount of bytes for the hash output from the PBKDF2 / scrypt hashing algorithm
	HashBytes int = 64
	// UniqueKey Key
	UniqueKey = "!!!!"
)

func HashPassword(rawPass, saltPassword string) string {
	hash := sha256.Sum256([]byte(rawPass + saltPassword))
	return hex.EncodeToString(hash[:])
}
