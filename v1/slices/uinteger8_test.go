package slices

import (
	"reflect"
	"testing"
)

func TestConvertUinteger8(t *testing.T) {
	i := []uint8{1, 2, 3, 4, 5}
	i = Uinteger8(i)
}

func TestUinteger8_ForEach(t *testing.T) {
	testCacheUinteger8s := []uint8{}

	type args struct {
		modifier func(index int, val uint8)
	}
	tests := []struct {
		name       string
		i          Uinteger8
		args       args
		cacheItems []uint8
	}{
		{
			name: "test uint8 slice for each",
			args: args{
				modifier: func(index int, val uint8) {
					testCacheUinteger8s = append(testCacheUinteger8s, val)
				},
			},
			i:          Uinteger8{1, 2, 3, 4, 5},
			cacheItems: []uint8{1, 2, 3, 4, 5},
		},
		{
			name: "test empty uint8 slice for each",
			args: args{
				modifier: func(index int, val uint8) {
					testCacheUinteger8s = append(testCacheUinteger8s, val)
				},
			},
			i:          Uinteger8{},
			cacheItems: []uint8{},
		},
	}
	for _, tt := range tests {
		testCacheUinteger8s = []uint8{}
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ForEach(tt.args.modifier)
		})
		if len(testCacheUinteger8s) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCacheUinteger8s),
			)
		}

		if !reflect.DeepEqual(testCacheUinteger8s, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCacheUinteger8s, tt.cacheItems)
		}
	}
}

func TestUinteger8_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val uint8) uint8
	}
	tests := []struct {
		name string
		i    Uinteger8
		args args
		want Uinteger8
	}{
		{
			name: "test success uint8 slice map",
			args: args{
				modifier: func(index int, val uint8) uint8 {
					return val + 1
				},
			},
			i:    Uinteger8{1, 2, 3, 4, 5},
			want: []uint8{2, 3, 4, 5, 6},
		},
		{
			name: "test empty uint8 slice map",
			args: args{
				modifier: func(index int, val uint8) uint8 {
					return 0
				},
			},
			i:    Uinteger8{},
			want: []uint8{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger8.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger8_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val uint8) bool
	}
	tests := []struct {
		name string
		i    Uinteger8
		args args
		want Uinteger8
	}{
		{
			name: "test uint8 slice filter",
			args: args{
				modifier: func(index int, val uint8) bool {
					return (val == 2)
				},
			},
			i:    Uinteger8{1, 2, 3, 4, 5},
			want: Uinteger8{2},
		},
		{
			name: "test empty uint8 slice filter",
			args: args{
				modifier: func(index int, val uint8) bool {
					return true
				},
			},
			i:    Uinteger8{},
			want: []uint8{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger8.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger8_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		i       Uinteger8
		args    args
		want    uint8
		want1   Uinteger8
		wantErr bool
	}{
		{
			name:    "test uint8 slice pop first",
			args:    args{index: 0},
			i:       Uinteger8{1, 2, 3, 4, 5},
			want:    1,
			want1:   Uinteger8{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "test uint8 slice pop last",
			args:    args{index: 4},
			i:       Uinteger8{1, 2, 3, 4, 5},
			want:    5,
			want1:   Uinteger8{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "test uint8 slice pop empty",
			args:    args{index: 0},
			i:       Uinteger8{},
			want:    0,
			want1:   Uinteger8{},
			wantErr: true,
		},
		{
			name:    "test uint8 slice pop out of bounds -1",
			args:    args{index: -1},
			i:       Uinteger8{1, 2, 3, 4, 5},
			want:    0,
			want1:   Uinteger8{},
			wantErr: true,
		},
		{
			name:    "test uint8 slice pop out of bounds 10",
			args:    args{index: 10},
			i:       Uinteger8{1, 2, 3, 4, 5},
			want:    0,
			want1:   Uinteger8{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.i.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uinteger8.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uinteger8.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Uinteger8.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUinteger8_Empty(t *testing.T) {
	tests := []struct {
		name string
		i    Uinteger8
		want bool
	}{
		{
			name: "test empty false",
			i:    []uint8{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "test empty true",
			i:    []uint8{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Empty(); got != tt.want {
				t.Errorf("Uinteger8.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger8_Contains(t *testing.T) {
	type args struct {
		value uint8
	}
	tests := []struct {
		name string
		i    Uinteger8
		args args
		want bool
	}{
		{
			name: "test does contain",
			i:    Uinteger8{1, 2, 3, 4, 5},
			args: args{value: 3},
			want: true,
		},
		{
			name: "test does not contain",
			i:    Uinteger8{1, 2, 3, 4, 5},
			args: args{value: 12},
			want: false,
		},
		{
			name: "test contain empty",
			i:    Uinteger8{},
			args: args{value: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Contains(tt.args.value); got != tt.want {
				t.Errorf("Uinteger8.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
