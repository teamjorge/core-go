package sets

import (
	"reflect"
	"sort"
	"testing"
)

func TestToUinteger32Slice(t *testing.T) {
	type args struct {
		in [][]uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "test single slice",
			args: args{in: [][]uint32{{1, 2, 3, 4, 5}}},
			want: []uint32{1, 2, 3, 4, 5},
		},
		{
			name: "test multiple slice",
			args: args{in: [][]uint32{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
			}},
			want: []uint32{1, 2, 3, 4, 5},
		},
		{
			name: "test empty array",
			args: args{in: [][]uint32{}},
			want: []uint32{},
		},
		{
			name: "test empty args",
			args: args{},
			want: []uint32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToUinteger32Slice(tt.args.in...)
			if len(got) != len(tt.want) {
				t.Errorf("Expected a result of length %d but received %d", len(tt.want), len(got))
			}
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUinteger32(t *testing.T) {
	type args struct {
		in [][]uint32
	}
	tests := []struct {
		name string
		args args
		want *Uinteger32
	}{
		{
			name: "test create new uint32 set simple",
			args: args{
				in: [][]uint32{{1, 2, 3, 4, 5}},
			},
			want: &Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test create new uint32 set empty",
			args: args{
				in: [][]uint32{{}},
			},
			want: &Uinteger32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUinteger32(tt.args.in...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUinteger32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger32_Add(t *testing.T) {
	type args struct {
		in []uint32
	}
	tests := []struct {
		name string
		i    Uinteger32
		args args
		want Uinteger32
	}{
		{
			name: "test add simple",
			args: args{
				in: []uint32{0},
			},
			i:    Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}, 0: struct{}{}},
		},
		{
			name: "test add empty",
			args: args{
				in: []uint32{},
			},
			i:    Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test add same",
			args: args{
				in: []uint32{1, 2, 3, 4, 5},
			},
			i:    Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
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

func TestUinteger32_Remove(t *testing.T) {
	type args struct {
		elem []uint32
	}
	tests := []struct {
		name string
		i    Uinteger32
		args args
		want Uinteger32
	}{
		{
			name: "test remove simple",
			args: args{
				elem: []uint32{0},
			},
			i:    Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}, 0: struct{}{}},
			want: Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test remove empty",
			args: args{
				elem: []uint32{0},
			},
			i:    Uinteger32{},
			want: Uinteger32{},
		},
		{
			name: "test remove nothing",
			args: args{
				elem: []uint32{},
			},
			i:    Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test remove not exists",
			args: args{
				elem: []uint32{0},
			},
			i:    Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
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

func TestUinteger32_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		i    Uinteger32
		want []uint32
	}{
		{
			name: "test to slice simple",
			i:    Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: []uint32{1, 2, 3, 4, 5},
		},
		{
			name: "test to slice empty",
			i:    Uinteger32{},
			want: []uint32{},
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

func TestUinteger32_Empty(t *testing.T) {
	tests := []struct {
		name string
		i    Uinteger32
		want bool
	}{
		{
			name: "test empty false",
			i:    Uinteger32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: false,
		},
		{
			name: "test empty true",
			i:    Uinteger32{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Empty(); got != tt.want {
				t.Errorf("Uinteger32.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
