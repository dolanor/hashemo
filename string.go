package hashemo

import (
	"encoding/hex"
	"fmt"
)

// FromHexString converts a classic hash output into emojis.
func FromHexString(s string) (string, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("hashemo: malformed hex input: %w", err)
	}

	return FromBytes(b), nil
}

// FromBytes converts a byte slice b into emojis.
func FromBytes(b []byte) string {
	s := make([]rune, 0, len(b))
	for _, octet := range b {
		s = append(s, emojis[octet])
	}
	return string(s)
}
