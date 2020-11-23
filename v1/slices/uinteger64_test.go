package slices

import (
	"reflect"
	"testing"
)

func TestConvertUinteger64(t *testing.T) {
	i := []uint64{1, 2, 3, 4, 5}
	i = Uinteger64(i)
}

func TestUinteger64_ForEach(t *testing.T) {
	testCacheUinteger64s := []uint64{}

	type args struct {
		modifier func(index int, val uint64)
	}
	tests := []struct {
		name       string
		i          Uinteger64
		args       args
		cacheItems []uint64
	}{
		{
			name: "test uint64 slice for each",
			args: args{
				modifier: func(index int, val uint64) {
					testCacheUinteger64s = append(testCacheUinteger64s, val)
				},
			},
			i:          Uinteger64{1, 2, 3, 4, 5},
			cacheItems: []uint64{1, 2, 3, 4, 5},
		},
		{
			name: "test empty uint64 slice for each",
			args: args{
				modifier: func(index int, val uint64) {
					testCacheUinteger64s = append(testCacheUinteger64s, val)
				},
			},
			i:          Uinteger64{},
			cacheItems: []uint64{},
		},
	}
	for _, tt := range tests {
		testCacheUinteger64s = []uint64{}
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ForEach(tt.args.modifier)
		})
		if len(testCacheUinteger64s) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCacheUinteger64s),
			)
		}

		if !reflect.DeepEqual(testCacheUinteger64s, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCacheUinteger64s, tt.cacheItems)
		}
	}
}

func TestUinteger64_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val uint64) uint64
	}
	tests := []struct {
		name string
		i    Uinteger64
		args args
		want Uinteger64
	}{
		{
			name: "test success uint64 slice map",
			args: args{
				modifier: func(index int, val uint64) uint64 {
					return val + 1
				},
			},
			i:    Uinteger64{1, 2, 3, 4, 5},
			want: []uint64{2, 3, 4, 5, 6},
		},
		{
			name: "test empty uint64 slice map",
			args: args{
				modifier: func(index int, val uint64) uint64 {
					return 0
				},
			},
			i:    Uinteger64{},
			want: []uint64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger64.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger64_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val uint64) bool
	}
	tests := []struct {
		name string
		i    Uinteger64
		args args
		want Uinteger64
	}{
		{
			name: "test uint64 slice filter",
			args: args{
				modifier: func(index int, val uint64) bool {
					return (val == 2)
				},
			},
			i:    Uinteger64{1, 2, 3, 4, 5},
			want: Uinteger64{2},
		},
		{
			name: "test empty uint64 slice filter",
			args: args{
				modifier: func(index int, val uint64) bool {
					return true
				},
			},
			i:    Uinteger64{},
			want: []uint64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger64.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger64_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		i       Uinteger64
		args    args
		want    uint64
		want1   Uinteger64
		wantErr bool
	}{
		{
			name:    "test uint64 slice pop first",
			args:    args{index: 0},
			i:       Uinteger64{1, 2, 3, 4, 5},
			want:    1,
			want1:   Uinteger64{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "test uint64 slice pop last",
			args:    args{index: 4},
			i:       Uinteger64{1, 2, 3, 4, 5},
			want:    5,
			want1:   Uinteger64{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "test uint64 slice pop empty",
			args:    args{index: 0},
			i:       Uinteger64{},
			want:    0,
			want1:   Uinteger64{},
			wantErr: true,
		},
		{
			name:    "test uint64 slice pop out of bounds -1",
			args:    args{index: -1},
			i:       Uinteger64{1, 2, 3, 4, 5},
			want:    0,
			want1:   Uinteger64{},
			wantErr: true,
		},
		{
			name:    "test uint64 slice pop out of bounds 10",
			args:    args{index: 10},
			i:       Uinteger64{1, 2, 3, 4, 5},
			want:    0,
			want1:   Uinteger64{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.i.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uinteger64.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uinteger64.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Uinteger64.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUinteger64_Empty(t *testing.T) {
	tests := []struct {
		name string
		i    Uinteger64
		want bool
	}{
		{
			name: "test empty false",
			i:    []uint64{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "test empty true",
			i:    []uint64{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Empty(); got != tt.want {
				t.Errorf("Uinteger64.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
