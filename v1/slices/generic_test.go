package slices

import "testing"

func TestIsEmpty(t *testing.T) {
	type args struct {
		s Slice
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test is empty string set",
			args: args{s: StringSlice{}},
			want: true,
		},
		{
			name: "test is not empty string set",
			args: args{s: StringSlice{"this", "is"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.s); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
