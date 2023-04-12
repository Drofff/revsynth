package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsInt(t *testing.T) {
	tests := []struct {
		title   string
		s       []int
		el      int
		wantRes bool
	}{
		{title: "empty s", s: []int{}, el: 2, wantRes: false},
		{title: "has el", s: []int{1, 2, 3, 4, 5}, el: 4, wantRes: true},
		{title: "does not have el", s: []int{1, 2, 3, 4, 5}, el: 9, wantRes: false},
	}
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			assert.Equal(t, test.wantRes, ContainsInt(test.s, test.el))
		})
	}
}
