package challenge1

import (
	"encoding/base64"
	"encoding/hex"
)

func Solve(src []byte) ([]byte, error) {
	buf := make([]byte, hex.DecodedLen(len(src)))
	bytes, err := hex.Decode(buf, src)
	if err != nil {
		return nil, err
	}

	l := base64.StdEncoding.EncodedLen(len(buf[:bytes]))
	out := make([]byte, l)
	base64.StdEncoding.Encode(out, buf[:bytes])

	return out, nil
}
