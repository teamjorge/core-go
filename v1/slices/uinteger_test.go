package slices

import (
	"reflect"
	"testing"
)

func TestConvertUinteger(t *testing.T) {
	i := []uint{1, 2, 3, 4, 5}
	i = Uinteger(i)
}

func TestUinteger_ForEach(t *testing.T) {
	testCacheUintegers := []uint{}

	type args struct {
		modifier func(index int, val uint)
	}
	tests := []struct {
		name       string
		i          Uinteger
		args       args
		cacheItems []uint
	}{
		{
			name: "test uint slice for each",
			args: args{
				modifier: func(index int, val uint) {
					testCacheUintegers = append(testCacheUintegers, val)
				},
			},
			i:          Uinteger{1, 2, 3, 4, 5},
			cacheItems: []uint{1, 2, 3, 4, 5},
		},
		{
			name: "test empty uint slice for each",
			args: args{
				modifier: func(index int, val uint) {
					testCacheUintegers = append(testCacheUintegers, val)
				},
			},
			i:          Uinteger{},
			cacheItems: []uint{},
		},
	}
	for _, tt := range tests {
		testCacheUintegers = []uint{}
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ForEach(tt.args.modifier)
		})
		if len(testCacheUintegers) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCacheUintegers),
			)
		}

		if !reflect.DeepEqual(testCacheUintegers, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCacheUintegers, tt.cacheItems)
		}
	}
}

func TestUinteger_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val uint) uint
	}
	tests := []struct {
		name string
		i    Uinteger
		args args
		want Uinteger
	}{
		{
			name: "test success uint slice map",
			args: args{
				modifier: func(index int, val uint) uint {
					return val + 1
				},
			},
			i:    Uinteger{1, 2, 3, 4, 5},
			want: []uint{2, 3, 4, 5, 6},
		},
		{
			name: "test empty uint slice map",
			args: args{
				modifier: func(index int, val uint) uint {
					return 0
				},
			},
			i:    Uinteger{},
			want: []uint{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val uint) bool
	}
	tests := []struct {
		name string
		i    Uinteger
		args args
		want Uinteger
	}{
		{
			name: "test uint slice filter",
			args: args{
				modifier: func(index int, val uint) bool {
					return (val == 2)
				},
			},
			i:    Uinteger{1, 2, 3, 4, 5},
			want: Uinteger{2},
		},
		{
			name: "test empty uint slice filter",
			args: args{
				modifier: func(index int, val uint) bool {
					return true
				},
			},
			i:    Uinteger{},
			want: []uint{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		i       Uinteger
		args    args
		want    uint
		want1   Uinteger
		wantErr bool
	}{
		{
			name:    "test uint slice pop first",
			args:    args{index: 0},
			i:       Uinteger{1, 2, 3, 4, 5},
			want:    1,
			want1:   Uinteger{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "test uint slice pop last",
			args:    args{index: 4},
			i:       Uinteger{1, 2, 3, 4, 5},
			want:    5,
			want1:   Uinteger{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "test uint slice pop empty",
			args:    args{index: 0},
			i:       Uinteger{},
			want:    0,
			want1:   Uinteger{},
			wantErr: true,
		},
		{
			name:    "test uint slice pop out of bounds -1",
			args:    args{index: -1},
			i:       Uinteger{1, 2, 3, 4, 5},
			want:    0,
			want1:   Uinteger{},
			wantErr: true,
		},
		{
			name:    "test uint slice pop out of bounds 10",
			args:    args{index: 10},
			i:       Uinteger{1, 2, 3, 4, 5},
			want:    0,
			want1:   Uinteger{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.i.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uinteger.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uinteger.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Uinteger.Pop() got1 = %v, want %v", got1, tt.want1)
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
			i:    []uint{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "test empty true",
			i:    []uint{},
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

func TestUinteger_Contains(t *testing.T) {
	type args struct {
		value uint
	}
	tests := []struct {
		name string
		i    Uinteger
		args args
		want bool
	}{
		{
			name: "test does contain",
			i:    Uinteger{1, 2, 3, 4, 5},
			args: args{value: 3},
			want: true,
		},
		{
			name: "test does not contain",
			i:    Uinteger{1, 2, 3, 4, 5},
			args: args{value: 12},
			want: false,
		},
		{
			name: "test contain empty",
			i:    Uinteger{},
			args: args{value: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Contains(tt.args.value); got != tt.want {
				t.Errorf("Uinteger.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
