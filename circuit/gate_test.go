package circuit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountControls(t *testing.T) {
	tests := []struct {
		title   string
		giveIn  []int
		wantOut int
	}{
		{title: "has controls", giveIn: []int{0, 0, 1, 1, 2, 2, 1, 0}, wantOut: 6},
		{title: "not controls", giveIn: []int{2, 2, 2, 2}, wantOut: 0},
		{title: "all controls", giveIn: []int{0, 1, 1, 0}, wantOut: 4},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			assert.Equal(t, test.wantOut, CountControls(test.giveIn))
		})
	}
}
