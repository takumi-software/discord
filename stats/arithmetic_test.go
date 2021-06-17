package stats

import "testing"

func Test_mean(t *testing.T) {
	type args struct {
		population []float64
	}
	tests := []struct {
		name     string
		args     args
		wantMean float64
	}{
		{
			name:     "If all values are zeroes, then the mean should be zero",
			args:     args{population: []float64{0, 0, 0, 0, 0}},
			wantMean: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMean := Mean(tt.args.population); gotMean != tt.wantMean {
				t.Errorf("mean() = %v, want %v", gotMean, tt.wantMean)
			}
		})
	}
}
