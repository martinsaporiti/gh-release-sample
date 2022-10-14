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
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 4},
		{2, 5},
	}

	for _, tc := range tests {
		got := next(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

}
