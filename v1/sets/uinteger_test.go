package sets

import (
	"reflect"
	"sort"
	"testing"
)

func TestToUintegerSlice(t *testing.T) {
	type args struct {
		in [][]uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "test single slice",
			args: args{in: [][]uint{{1, 2, 3, 4, 5}}},
			want: []uint{1, 2, 3, 4, 5},
		},
		{
			name: "test multiple slice",
			args: args{in: [][]uint{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
			}},
			want: []uint{1, 2, 3, 4, 5},
		},
		{
			name: "test empty array",
			args: args{in: [][]uint{}},
			want: []uint{},
		},
		{
			name: "test empty args",
			args: args{},
			want: []uint{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToUintegerSlice(tt.args.in...)
			if len(got) != len(tt.want) {
				t.Errorf("Expected a result of length %d but received %d", len(tt.want), len(got))
			}
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUinteger(t *testing.T) {
	type args struct {
		in [][]uint
	}
	tests := []struct {
		name string
		args args
		want *Uinteger
	}{
		{
			name: "test create new uint set simple",
			args: args{
				in: [][]uint{{1, 2, 3, 4, 5}},
			},
			want: &Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test create new uint set empty",
			args: args{
				in: [][]uint{{}},
			},
			want: &Uinteger{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUinteger(tt.args.in...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUinteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger_Add(t *testing.T) {
	type args struct {
		in []uint
	}
	tests := []struct {
		name string
		i    Uinteger
		args args
		want Uinteger
	}{
		{
			name: "test add simple",
			args: args{
				in: []uint{0},
			},
			i:    Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}, 0: struct{}{}},
		},
		{
			name: "test add empty",
			args: args{
				in: []uint{},
			},
			i:    Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test add same",
			args: args{
				in: []uint{1, 2, 3, 4, 5},
			},
			i:    Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.i.Add(tt.args.in...)
			if !reflect.DeepEqual(tt.i, tt.want) {
				t.Errorf("Add() = %v, want %v", tt.i, tt.want)
			}
		})
	}
}

func TestUinteger_Remove(t *testing.T) {
	type args struct {
		elem []uint
	}
	tests := []struct {
		name string
		i    Uinteger
		args args
		want Uinteger
	}{
		{
			name: "test remove simple",
			args: args{
				elem: []uint{0},
			},
			i:    Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}, 0: struct{}{}},
			want: Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test remove empty",
			args: args{
				elem: []uint{0},
			},
			i:    Uinteger{},
			want: Uinteger{},
		},
		{
			name: "test remove nothing",
			args: args{
				elem: []uint{},
			},
			i:    Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test remove not exists",
			args: args{
				elem: []uint{0},
			},
			i:    Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.i.Remove(tt.args.elem...)
			if !reflect.DeepEqual(tt.i, tt.want) {
				t.Errorf("Remove() = %v, want %v", tt.i, tt.want)
			}
		})
	}
}

func TestUinteger_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		i    Uinteger
		want []uint
	}{
		{
			name: "test to slice simple",
			i:    Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: []uint{1, 2, 3, 4, 5},
		},
		{
			name: "test to slice empty",
			i:    Uinteger{},
			want: []uint{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.i.ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Expected a result of length %d but received %d", len(tt.want), len(got))
			}
			// If sort doesn't work, add custom sorting
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger_Empty(t *testing.T) {
	tests := []struct {
		name string
		i    Uinteger
		want bool
	}{
		{
			name: "test empty false",
			i:    Uinteger{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: false,
		},
		{
			name: "test empty true",
			i:    Uinteger{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Empty(); got != tt.want {
				t.Errorf("Uinteger.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
