package util

import "encoding/base64"

func Hex2Base64(input []byte) []byte {
	len := base64.StdEncoding.EncodedLen(len(input))
	buffer := make([]byte, len)

	base64.StdEncoding.Encode(buffer, input)

	return buffer
}
