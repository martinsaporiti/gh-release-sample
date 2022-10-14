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
		{10, 11},
		{14, 15},
		{15, 16},
		{15, 17},
		{17, 18},
		{18, 19},
		{19, 20},
		{20, 21},
		{21, 22},
		{22, 23},
		{23, 24},
	}

	for _, tc := range tests {
		got := next(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

}
