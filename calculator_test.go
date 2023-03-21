package rpncalculator

import "testing"

func TestCalculator(t *testing.T) {
	tests := []struct {
		expr    string
		want    float64
		wantErr bool
	}{
		{
			expr: "3 4 5 * +",
			want: 23,
		},
		{
			expr: "3 4 5 * -",
			want: -17,
		},
		{
			expr: "10 20 -",
			want: -10,
		},
		{
			expr: "3 5 6 1 + - +",
			want: 1,
		},
		{
			expr:    "3 5 6 + %",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		got, err := Evaluate(tt.expr)
		if tt.wantErr && err == nil {
			t.Fatalf("expected error got: %q", err)
		}

		if got != tt.want {
			t.Fatalf("expected: %v got: %v", tt.want, got)
		}
	}
}
