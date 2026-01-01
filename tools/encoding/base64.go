package encoding

import (
	"encoding/base64"
)


func ToBase64(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func FromBase64(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}