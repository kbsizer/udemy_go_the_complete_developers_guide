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

// TODO: Use this to write tests for SAVE and LOAD methods
// https://www.toptal.com/go/your-introductory-course-to-testing-with-go

/*
maybe put this into a markdown doc or something?

https://blog.golang.org/slices

new() --- built-in function;

make(slice,initialLength[],capacity]) -- built-in function; allocates a new array and creates a slice header to describe it,
		  all at once. The make function takes three arguments: the type of the slice, its
		  initial length, and its capacity, which is the length of the array that make allocates
		  to hold the slice data. (if capacity is omitted, capacity=length)

		  copy(targetSlice,sourceSlice) -- built-in function; copies data from sourceSlice to targetSlice;
will only copy the minimum of the lengths of the two given slices; returns the number of elements copied

append(slice, element) -- built-in function; returns the result of appending element to slice; allocates more memory if necessary
append(slice1, slice2...) -- using the "variadic" or "ellipsis" operator (See: https://yourbasic.org/golang/three-dots-ellipsis/)
*/
