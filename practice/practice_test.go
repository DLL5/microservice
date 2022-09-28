package practice

import "testing"

func Test_chanPractice(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"one"}, {"two"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chanPractice()
		})
	}
}
