package chars

import "testing"

func TestChar_Reverse(t *testing.T) {
	tests := []struct {
		name string
		c    Char
		want string
	}{
		{
			name: "test reverse string",
			c:    "this is my sentence",
			want: "ecnetnes ym si siht",
		},
		{
			name: "test reverse string more chars",
			c:    "abcdefg11!!!||||fdsasd!!!@@@@//&$^%#@~!",
			want: "!~@#%^$&//@@@@!!!dsasdf||||!!!11gfedcba",
		},
		{
			name: "test reverse string emptys",
			c:    "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Reverse(); got != tt.want {
				t.Errorf("Char.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
