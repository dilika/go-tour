package main

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a       int
		b       int
		want    int
		wantErr bool
	}{
		{name: "exact division", a: 20, b: 5, want: 4},
		{name: "integer truncation", a: 7, b: 2, want: 3},
		{name: "negative result", a: -9, b: 2, want: -4},
		{name: "zero numerator", a: 0, b: 5, want: 0},
		{name: "zero denominator", a: 5, b: 0, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.a, tt.b)

			if (err != nil) != tt.wantErr {
				t.Fatalf("divide(%d, %d) error = %v, wantErr %v",
					tt.a, tt.b, err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("divide(%d, %d) = %d, want %d",
					tt.a, tt.b, got, tt.want)
			}
		})
	}
}
