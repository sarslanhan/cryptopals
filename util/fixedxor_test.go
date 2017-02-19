package util

import (
	"encoding/hex"
	"testing"
)

type FixedXorTestCase struct {
	a        string
	b        string
	expected string
}

var fixedXorTestCases = [...]FixedXorTestCase{
	FixedXorTestCase{
		"1c0111001f010100061a024b53535009181c",
		"686974207468652062756c6c277320657965",
		"746865206b696420646f6e277420706c6179",
	},
}

func TestFixedXor(t *testing.T) {
	for _, tc := range fixedXorTestCases {
		a, _ := hex.DecodeString(tc.a)
		b, _ := hex.DecodeString(tc.b)

		got := hex.EncodeToString(FixedXor(a, b))

		if got != tc.expected {
			t.Errorf("Expected: '%s' Got: '%s'", tc.expected, got)
		}
	}
}
