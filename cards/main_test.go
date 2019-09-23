package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// silly overkill to use table-based testing idiom here, but... for consistency...
		{"Executing main()"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
