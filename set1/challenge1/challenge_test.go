package challenge1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		hex  string
		want string
	}{
		{
			"hex to base64",
			"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			result, err := Solve([]byte(tt.hex))
			assert.NoError(err)

			//encoded := base64.StdEncoding.EncodeToString(result)

			assert.Equal(tt.want, string(result))
		})
	}
}
