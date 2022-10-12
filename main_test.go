package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	type test struct {
		input int
		want  int
	}

	tests := []test{
		{1, 2},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 8},
		{8, 9},
		{9, 10},
	}

	for _, tc := range tests {
		got := next(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

}
