package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string // test name only (no input params or return values)
	}{
		// silly overkill to use table-based testing idiom here, but... for consistency...
		// Syntax: { "test name / description", input parameter value(s), expected return value(s) }
		{"Executing main()"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_estPi(t *testing.T) {
	tests := []struct {
		name string  // test name
		want float64 // return value (no input parameters)
	}{
		// Syntax: { "test name / description", input parameter value(s), expected return value(s) }
		{"Executing estPi()", 3.14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := estPi(); got != tt.want {
				t.Errorf("estPi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hypotenuse(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string  // test name
		args args    // input parameters (2)
		want float64 // return value
	}{
		{"3-4-5 test", args{3.0, 4.0}, 5.0},
		{"6-8-10 test", args{6.0, 8.0}, 10.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hypotenuse(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("hypotenuse() = %v, want %v", got, tt.want)
			}
		})
	}
}
