package sets

import (
	"reflect"
	"sort"
	"testing"
)

func TestString(t *testing.T) {
	type args struct {
		in [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test single slice",
			args: args{in: [][]string{{"this", "is", "unique", "this", "is", "distinct"}}},
			want: []string{"distinct", "is", "this", "unique"},
		},
		{
			name: "test multiple slice",
			args: args{in: [][]string{
				{"values", "from", "first", "slice"},
				{"values", "from", "second", "slice"},
			}},
			want: []string{"first", "from", "second", "slice", "values"},
		},
		{
			name: "test empty array",
			args: args{in: [][]string{}},
			want: []string{},
		},
		{
			name: "test empty args",
			args: args{},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := String(tt.args.in...)
			if len(got) != len(tt.want) {
				t.Errorf("Expected a result of length %d but received %d", len(tt.want), len(got))
			}
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
