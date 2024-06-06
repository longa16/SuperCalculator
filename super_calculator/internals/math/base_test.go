package math

import (
	"testing"
)

func TestAddition(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"simple addition", args{4, 5}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Addition(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Addition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"simple multiplication", args{4, 5}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiplication(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Multiplication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtraction(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"simple multiplication", args{4, 5}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtraction(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Subtraction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"simple division", args{4, 5}, 0.8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Division(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Division() = %v, want %v", got, tt.want)
			}
		})
	}
}
