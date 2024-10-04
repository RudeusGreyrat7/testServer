package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) ([]byte, error) {
	data := make([]byte, n)
	nRead, err := rand.Read(data)
	if err != nil {
		return nil, fmt.Errorf("bytes: %v", err)
	}
	if nRead < n {
		return nil, fmt.Errorf("bytes: didn't read enough random bytes")
	}
	return data, nil
}

func String(n int) (string, error) {
	data, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("string: %v", err)
	}
	return base64.URLEncoding.EncodeToString(data), nil
}
