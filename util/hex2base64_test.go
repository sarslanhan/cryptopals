package util

import (
	"encoding/hex"
	"testing"
)

type Hex2Base64ConversionTestCase struct {
	inputString  string
	outputString string
}

var hex2Base64ConversionTestCases = [...]Hex2Base64ConversionTestCase{
	Hex2Base64ConversionTestCase{
		"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
		"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
	},
}

func TestHex2Base64(t *testing.T) {

	for _, tc := range hex2Base64ConversionTestCases {
		input, _ := hex.DecodeString(tc.inputString)

		output := Hex2Base64(input)

		outputString := string(output[:])

		if outputString != tc.outputString {
			t.Errorf("Expected: '%s' Got: '%s'", tc.outputString, outputString)
		}
	}

}
