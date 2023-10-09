package hashemo

import (
	"encoding/hex"
	"strings"
)

// FromHexString converts a classic hash output into emojis.
func FromHexString(s string) string {
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		var i rune
		switch {
		case r >= '0' && r <= '9':
			i = r - '0'
		case r >= 'a' && r <= 'f':
			i = r - 'a'
		default:
			return r
		}
		return rune(emojis[i])
	}, s)

	return s
}

// FromBytes converts a byte slice b into emojis.
func FromBytes(b []byte) string {
	s := hex.EncodeToString(b)
	return FromHexString(s)
}
