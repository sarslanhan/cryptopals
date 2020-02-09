package challenge2

import (
	"encoding/hex"
)

func Solve(lhs, rhs []byte) ([]byte, error) {
	lbuf := make([]byte, hex.DecodedLen(len(lhs)))
	_, err := hex.Decode(lbuf, lhs)
	if err != nil {
		return nil, err
	}
	rbuf := make([]byte, hex.DecodedLen(len(rhs)))
	_, err = hex.Decode(rbuf, rhs)
	if err != nil {
		return nil, err
	}

	for i := range lbuf {
		lbuf[i] ^= rbuf[i]
	}

	rbuf = make([]byte, hex.EncodedLen(len(lbuf)))

	hex.Encode(rbuf, lbuf)

	return rbuf, nil
}
