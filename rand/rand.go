package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) ([]byte, error) {
  b := make([]byte, n)
  nRead, err := rand.Read(b)
  if err!= nil {
    return nil, fmt.Errorf("failed to generate random bytes: %v", err)
  }
  if nRead < n {
    return nil, fmt.Errorf("not enough random bytes: %v", err)
  }
  return b, nil
}

// String returns a randon string using crypto/rand.
// n is the number of bytes being used to generate the random string.
func String(n int) (string, error) {
  b, err := Bytes(n)
  if err!= nil {
    return "", fmt.Errorf("string: %w", err)
  }
  return base64.URLEncoding.EncodeToString(b), nil
}