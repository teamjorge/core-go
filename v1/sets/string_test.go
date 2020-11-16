package sets

import (
	"reflect"
	"sort"
	"testing"
)

func TestToStringSlice(t *testing.T) {
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
			got := ToStringSlice(tt.args.in...)
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

func TestNewString(t *testing.T) {
	type args struct {
		in [][]string
	}
	tests := []struct {
		name string
		args args
		want *String
	}{
		{
			name: "test create new string set simple",
			args: args{
				in: [][]string{{"this", "this", "is", "is"}},
			},
			want: &String{
				"this": struct{}{},
				"is":   struct{}{},
			},
		},
		{
			name: "test create new string set empty",
			args: args{
				in: [][]string{{}},
			},
			want: &String{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewString(tt.args.in...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Add(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		s    String
		args args
		want String
	}{
		{
			name: "test add simple",
			args: args{
				in: []string{"new"},
			},
			s: String{"this": struct{}{}, "is": struct{}{}},
			want: String{
				"this": struct{}{},
				"is":   struct{}{},
				"new":  struct{}{},
			},
		},
		{
			name: "test add empty",
			args: args{
				in: []string{},
			},
			s: String{"this": struct{}{}, "is": struct{}{}},
			want: String{
				"this": struct{}{},
				"is":   struct{}{},
			},
		},
		{
			name: "test add same",
			args: args{
				in: []string{"this", "is"},
			},
			s: String{"this": struct{}{}, "is": struct{}{}},
			want: String{
				"this": struct{}{},
				"is":   struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.in...)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Add() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestString_Remove(t *testing.T) {
	type args struct {
		elem []string
	}
	tests := []struct {
		name string
		s    String
		args args
		want String
	}{
		{
			name: "test remove simple",
			args: args{
				elem: []string{"is"},
			},
			s: String{"this": struct{}{}, "is": struct{}{}},
			want: String{
				"this": struct{}{},
			},
		},
		{
			name: "test remove empty",
			args: args{
				elem: []string{},
			},
			s: String{"this": struct{}{}, "is": struct{}{}},
			want: String{
				"this": struct{}{},
				"is":   struct{}{},
			},
		},
		{
			name: "test remove not exists",
			args: args{
				elem: []string{"new"},
			},
			s: String{"this": struct{}{}, "is": struct{}{}},
			want: String{
				"this": struct{}{},
				"is":   struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.elem...)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Remove() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestString_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		s    String
		want []string
	}{
		{
			name: "test to slice simple",
			s:    String{"this": struct{}{}, "is": struct{}{}},
			want: []string{"is", "this"},
		},
		{
			name: "test to slice empty",
			s:    String{},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Expected a result of length %d but received %d", len(tt.want), len(got))
			}
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Empty(t *testing.T) {
	tests := []struct {
		name string
		s    String
		want bool
	}{
		{
			name: "test empty true",
			s:    String{"this": struct{}{}, "is": struct{}{}},
			want: false,
		},
		{
			name: "test empty false",
			s:    String{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Empty(); got != tt.want {
				t.Errorf("String.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
