package slices

import (
	"reflect"
	"testing"
)

func TestConvertUinteger32(t *testing.T) {
	i := []uint32{1, 2, 3, 4, 5}
	i = Uinteger32(i)
}

func TestUinteger32_ForEach(t *testing.T) {
	testCacheUinteger32s := []uint32{}

	type args struct {
		modifier func(index int, val uint32)
	}
	tests := []struct {
		name       string
		i          Uinteger32
		args       args
		cacheItems []uint32
	}{
		{
			name: "test uint32 slice for each",
			args: args{
				modifier: func(index int, val uint32) {
					testCacheUinteger32s = append(testCacheUinteger32s, val)
				},
			},
			i:          Uinteger32{1, 2, 3, 4, 5},
			cacheItems: []uint32{1, 2, 3, 4, 5},
		},
		{
			name: "test empty uint32 slice for each",
			args: args{
				modifier: func(index int, val uint32) {
					testCacheUinteger32s = append(testCacheUinteger32s, val)
				},
			},
			i:          Uinteger32{},
			cacheItems: []uint32{},
		},
	}
	for _, tt := range tests {
		testCacheUinteger32s = []uint32{}
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ForEach(tt.args.modifier)
		})
		if len(testCacheUinteger32s) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCacheUinteger32s),
			)
		}

		if !reflect.DeepEqual(testCacheUinteger32s, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCacheUinteger32s, tt.cacheItems)
		}
	}
}

func TestUinteger32_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val uint32) uint32
	}
	tests := []struct {
		name string
		i    Uinteger32
		args args
		want Uinteger32
	}{
		{
			name: "test success uint32 slice map",
			args: args{
				modifier: func(index int, val uint32) uint32 {
					return val + 1
				},
			},
			i:    Uinteger32{1, 2, 3, 4, 5},
			want: []uint32{2, 3, 4, 5, 6},
		},
		{
			name: "test empty uint32 slice map",
			args: args{
				modifier: func(index int, val uint32) uint32 {
					return 0
				},
			},
			i:    Uinteger32{},
			want: []uint32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger32.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger32_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val uint32) bool
	}
	tests := []struct {
		name string
		i    Uinteger32
		args args
		want Uinteger32
	}{
		{
			name: "test uint32 slice filter",
			args: args{
				modifier: func(index int, val uint32) bool {
					return (val == 2)
				},
			},
			i:    Uinteger32{1, 2, 3, 4, 5},
			want: Uinteger32{2},
		},
		{
			name: "test empty uint32 slice filter",
			args: args{
				modifier: func(index int, val uint32) bool {
					return true
				},
			},
			i:    Uinteger32{},
			want: []uint32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger32.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger32_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		i       Uinteger32
		args    args
		want    uint32
		want1   Uinteger32
		wantErr bool
	}{
		{
			name:    "test uint32 slice pop first",
			args:    args{index: 0},
			i:       Uinteger32{1, 2, 3, 4, 5},
			want:    1,
			want1:   Uinteger32{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "test uint32 slice pop last",
			args:    args{index: 4},
			i:       Uinteger32{1, 2, 3, 4, 5},
			want:    5,
			want1:   Uinteger32{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "test uint32 slice pop empty",
			args:    args{index: 0},
			i:       Uinteger32{},
			want:    0,
			want1:   Uinteger32{},
			wantErr: true,
		},
		{
			name:    "test uint32 slice pop out of bounds -1",
			args:    args{index: -1},
			i:       Uinteger32{1, 2, 3, 4, 5},
			want:    0,
			want1:   Uinteger32{},
			wantErr: true,
		},
		{
			name:    "test uint32 slice pop out of bounds 10",
			args:    args{index: 10},
			i:       Uinteger32{1, 2, 3, 4, 5},
			want:    0,
			want1:   Uinteger32{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.i.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uinteger32.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uinteger32.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Uinteger32.Pop() got1 = %v, want %v", got1, tt.want1)
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
			i:    []uint32{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "test empty true",
			i:    []uint32{},
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

func TestUinteger32_Contains(t *testing.T) {
	type args struct {
		value uint32
	}
	tests := []struct {
		name string
		i    Uinteger32
		args args
		want bool
	}{
		{
			name: "test does contain",
			i:    Uinteger32{1, 2, 3, 4, 5},
			args: args{value: 3},
			want: true,
		},
		{
			name: "test does not contain",
			i:    Uinteger32{1, 2, 3, 4, 5},
			args: args{value: 12},
			want: false,
		},
		{
			name: "test contain empty",
			i:    Uinteger32{},
			args: args{value: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Contains(tt.args.value); got != tt.want {
				t.Errorf("Uinteger32.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
