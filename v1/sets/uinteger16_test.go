package sets

import (
	"reflect"
	"sort"
	"testing"
)

func TestToUinteger16Slice(t *testing.T) {
	type args struct {
		in [][]uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "test single slice",
			args: args{in: [][]uint16{{1, 2, 3, 4, 5}}},
			want: []uint16{1, 2, 3, 4, 5},
		},
		{
			name: "test multiple slice",
			args: args{in: [][]uint16{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
			}},
			want: []uint16{1, 2, 3, 4, 5},
		},
		{
			name: "test empty array",
			args: args{in: [][]uint16{}},
			want: []uint16{},
		},
		{
			name: "test empty args",
			args: args{},
			want: []uint16{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToUinteger16Slice(tt.args.in...)
			if len(got) != len(tt.want) {
				t.Errorf("Expected a result of length %d but received %d", len(tt.want), len(got))
			}
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUinteger16(t *testing.T) {
	type args struct {
		in [][]uint16
	}
	tests := []struct {
		name string
		args args
		want *Uinteger16
	}{
		{
			name: "test create new uint16 set simple",
			args: args{
				in: [][]uint16{{1, 2, 3, 4, 5}},
			},
			want: &Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test create new uint16 set empty",
			args: args{
				in: [][]uint16{{}},
			},
			want: &Uinteger16{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUinteger16(tt.args.in...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUinteger16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger16_Add(t *testing.T) {
	type args struct {
		in []uint16
	}
	tests := []struct {
		name string
		i    Uinteger16
		args args
		want Uinteger16
	}{
		{
			name: "test add simple",
			args: args{
				in: []uint16{0},
			},
			i:    Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}, 0: struct{}{}},
		},
		{
			name: "test add empty",
			args: args{
				in: []uint16{},
			},
			i:    Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test add same",
			args: args{
				in: []uint16{1, 2, 3, 4, 5},
			},
			i:    Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
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

func TestUinteger16_Remove(t *testing.T) {
	type args struct {
		elem []uint16
	}
	tests := []struct {
		name string
		i    Uinteger16
		args args
		want Uinteger16
	}{
		{
			name: "test remove simple",
			args: args{
				elem: []uint16{0},
			},
			i:    Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}, 0: struct{}{}},
			want: Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test remove empty",
			args: args{
				elem: []uint16{0},
			},
			i:    Uinteger16{},
			want: Uinteger16{},
		},
		{
			name: "test remove nothing",
			args: args{
				elem: []uint16{},
			},
			i:    Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test remove not exists",
			args: args{
				elem: []uint16{0},
			},
			i:    Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
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

func TestUinteger16_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		i    Uinteger16
		want []uint16
	}{
		{
			name: "test to slice simple",
			i:    Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: []uint16{1, 2, 3, 4, 5},
		},
		{
			name: "test to slice empty",
			i:    Uinteger16{},
			want: []uint16{},
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

func TestUinteger16_Empty(t *testing.T) {
	tests := []struct {
		name string
		i    Uinteger16
		want bool
	}{
		{
			name: "test empty false",
			i:    Uinteger16{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: false,
		},
		{
			name: "test empty true",
			i:    Uinteger16{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Empty(); got != tt.want {
				t.Errorf("Uinteger16.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
