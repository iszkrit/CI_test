package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		num []int
	}
	type testdata struct {
		name string
		args args
		want []string
	}
	tests := []testdata{
		{
			name: "UnitTest1",
			args: args{
				num: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			},
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"},
		},
		{
			name: "UnitTest2",
			args: args{
				num: []int{0, 15, 30, 45, 100},
			},
			want: []string{"FizzBuzz", "FizzBuzz", "FizzBuzz", "FizzBuzz", "Buzz"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}