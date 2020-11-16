package numberutil

import (
	"testing"
)

func TestToFixed(t *testing.T) {
	type args struct {
		num       float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"0.045", args{0.045, 4}, 0.045},
		{"0.045", args{0.045, 3}, 0.045},
		{"0.045", args{0.045, 2}, 0.05},
		{"0.045", args{0.045, 1}, 0.0},

		{"0.12", args{0.12, 4}, 0.12},
		{"0.12", args{0.12, 3}, 0.12},
		{"0.12", args{0.12, 2}, 0.12},
		{"0.12", args{0.12, 1}, 0.1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFixed(tt.args.num, tt.args.precision); got != tt.want {
				t.Errorf("ToFixed() = %v, want %v", got, tt.want)
			}
		})
	}
}
