package sets

import "testing"

func TestIsEmpty(t *testing.T) {
	type args struct {
		s Set
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test is empty string set",
			args: args{s: NewString()},
			want: true,
		},
		{
			name: "test is not empty string set",
			args: args{s: NewString([]string{"this", "is"})},
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
