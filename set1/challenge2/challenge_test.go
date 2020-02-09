package challenge2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	type args struct {
		lhs, rhs string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"fixed or",
			args{
				"1c0111001f010100061a024b53535009181c",
				"686974207468652062756c6c277320657965",
			},
			"746865206b696420646f6e277420706c6179",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			result, err := Solve([]byte(tt.args.lhs), []byte(tt.args.rhs))
			assert.NoError(err)

			assert.Equal(tt.want, string(result))
		})
	}
}
