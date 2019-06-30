package main

import (
	"reflect"
	"testing"
)

func TestPatternSearch(t *testing.T) {
	var tests = []struct {
		name         string
		pattern      string
		textToSearch string
		want         []int
	}{
		{"Successful when pattern is all upper case", "PETER", "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", []int{1, 20, 75}},
		{"Successful when pattern is all lower case", "peter", "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", []int{1, 20, 75}},
		{"Successful when pattern is a mix of upper and lower case", "peTer", "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", []int{1, 20, 75}},
		{"Successfull pattern match", "pi", "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", []int{30, 37, 43, 51, 58}},
		{"Return empty if no match", "Peterz", "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", nil},
		{"Return empty if pattern is longer than text to search", "Peterskjdhfs", "Pet", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PatternSearch(tt.textToSearch, tt.pattern)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PatternSearch:(%s) got %+v, want %+v", tt.pattern, got, tt.want)
			}
		})

	}
}

func TestAbs(t *testing.T) {
	data := -12
	want := 12
	got := Abs(data)
	if got != want {
		t.Errorf("Abs:(%d) got %d, want %d", data, got, 12)
	}
}
