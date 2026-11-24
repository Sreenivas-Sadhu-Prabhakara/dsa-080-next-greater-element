package main

import (
	"reflect"
	"testing"
)

func TestNextGreater(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{6, 8, 0, 1, 3}, []int{8, -1, 1, 3, -1}},
		{[]int{1, 2, 3}, []int{2, 3, -1}},
		{[]int{3, 2, 1}, []int{-1, -1, -1}},
		{[]int{2, 2}, []int{-1, -1}}, // strictly greater, equals don't count
		{[]int{}, []int{}},
		{[]int{7}, []int{-1}},
	}
	for _, c := range cases {
		if got := NextGreater(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("NextGreater(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
