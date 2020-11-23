package sets

import (
	"reflect"
	"sort"
	"testing"
)

func TestToInteger32Slice(t *testing.T) {
	type args struct {
		in [][]int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "test single slice",
			args: args{in: [][]int32{{1, 2, 3, 4, 5}}},
			want: []int32{1, 2, 3, 4, 5},
		},
		{
			name: "test multiple slice",
			args: args{in: [][]int32{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
			}},
			want: []int32{1, 2, 3, 4, 5},
		},
		{
			name: "test empty array",
			args: args{in: [][]int32{}},
			want: []int32{},
		},
		{
			name: "test empty args",
			args: args{},
			want: []int32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToInteger32Slice(tt.args.in...)
			if len(got) != len(tt.want) {
				t.Errorf("Expected a result of length %d but received %d", len(tt.want), len(got))
			}
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integer32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInteger32(t *testing.T) {
	type args struct {
		in [][]int32
	}
	tests := []struct {
		name string
		args args
		want *Integer32
	}{
		{
			name: "test create new int32 set simple",
			args: args{
				in: [][]int32{{1, 2, 3, 4, 5}},
			},
			want: &Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test create new int32 set empty",
			args: args{
				in: [][]int32{{}},
			},
			want: &Integer32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInteger32(tt.args.in...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInteger32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger32_Add(t *testing.T) {
	type args struct {
		in []int32
	}
	tests := []struct {
		name string
		i    Integer32
		args args
		want Integer32
	}{
		{
			name: "test add simple",
			args: args{
				in: []int32{0},
			},
			i:    Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}, 0: struct{}{}},
		},
		{
			name: "test add empty",
			args: args{
				in: []int32{},
			},
			i:    Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test add same",
			args: args{
				in: []int32{1, 2, 3, 4, 5},
			},
			i:    Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
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

func TestInteger32_Remove(t *testing.T) {
	type args struct {
		elem []int32
	}
	tests := []struct {
		name string
		i    Integer32
		args args
		want Integer32
	}{
		{
			name: "test remove simple",
			args: args{
				elem: []int32{0},
			},
			i:    Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}, 0: struct{}{}},
			want: Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test remove empty",
			args: args{
				elem: []int32{0},
			},
			i:    Integer32{},
			want: Integer32{},
		},
		{
			name: "test remove nothing",
			args: args{
				elem: []int32{},
			},
			i:    Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		},
		{
			name: "test remove not exists",
			args: args{
				elem: []int32{0},
			},
			i:    Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
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

func TestInteger32_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		i    Integer32
		want []int32
	}{
		{
			name: "test to slice simple",
			i:    Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: []int32{1, 2, 3, 4, 5},
		},
		{
			name: "test to slice empty",
			i:    Integer32{},
			want: []int32{},
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

func TestInteger32_Empty(t *testing.T) {
	tests := []struct {
		name string
		i    Integer32
		want bool
	}{
		{
			name: "test empty false",
			i:    Integer32{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}},
			want: false,
		},
		{
			name: "test empty true",
			i:    Integer32{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Empty(); got != tt.want {
				t.Errorf("Integer32.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
