package slices

import (
	"reflect"
	"testing"
)

func TestConvertUinteger16(t *testing.T) {
	i := []uint16{1, 2, 3, 4, 5}
	i = Uinteger16(i)
}

func TestUinteger16_ForEach(t *testing.T) {
	testCacheUinteger16s := []uint16{}

	type args struct {
		modifier func(index int, val uint16)
	}
	tests := []struct {
		name       string
		i          Uinteger16
		args       args
		cacheItems []uint16
	}{
		{
			name: "test uint16 slice for each",
			args: args{
				modifier: func(index int, val uint16) {
					testCacheUinteger16s = append(testCacheUinteger16s, val)
				},
			},
			i:          Uinteger16{1, 2, 3, 4, 5},
			cacheItems: []uint16{1, 2, 3, 4, 5},
		},
		{
			name: "test empty uint16 slice for each",
			args: args{
				modifier: func(index int, val uint16) {
					testCacheUinteger16s = append(testCacheUinteger16s, val)
				},
			},
			i:          Uinteger16{},
			cacheItems: []uint16{},
		},
	}
	for _, tt := range tests {
		testCacheUinteger16s = []uint16{}
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ForEach(tt.args.modifier)
		})
		if len(testCacheUinteger16s) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCacheUinteger16s),
			)
		}

		if !reflect.DeepEqual(testCacheUinteger16s, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCacheUinteger16s, tt.cacheItems)
		}
	}
}

func TestUinteger16_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val uint16) uint16
	}
	tests := []struct {
		name string
		i    Uinteger16
		args args
		want Uinteger16
	}{
		{
			name: "test success uint16 slice map",
			args: args{
				modifier: func(index int, val uint16) uint16 {
					return val + 1
				},
			},
			i:    Uinteger16{1, 2, 3, 4, 5},
			want: []uint16{2, 3, 4, 5, 6},
		},
		{
			name: "test empty uint16 slice map",
			args: args{
				modifier: func(index int, val uint16) uint16 {
					return 0
				},
			},
			i:    Uinteger16{},
			want: []uint16{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger16.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger16_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val uint16) bool
	}
	tests := []struct {
		name string
		i    Uinteger16
		args args
		want Uinteger16
	}{
		{
			name: "test uint16 slice filter",
			args: args{
				modifier: func(index int, val uint16) bool {
					return (val == 2)
				},
			},
			i:    Uinteger16{1, 2, 3, 4, 5},
			want: Uinteger16{2},
		},
		{
			name: "test empty uint16 slice filter",
			args: args{
				modifier: func(index int, val uint16) bool {
					return true
				},
			},
			i:    Uinteger16{},
			want: []uint16{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uinteger16.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger16_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		i       Uinteger16
		args    args
		want    uint16
		want1   Uinteger16
		wantErr bool
	}{
		{
			name:    "test uint16 slice pop first",
			args:    args{index: 0},
			i:       Uinteger16{1, 2, 3, 4, 5},
			want:    1,
			want1:   Uinteger16{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "test uint16 slice pop last",
			args:    args{index: 4},
			i:       Uinteger16{1, 2, 3, 4, 5},
			want:    5,
			want1:   Uinteger16{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "test uint16 slice pop empty",
			args:    args{index: 0},
			i:       Uinteger16{},
			want:    0,
			want1:   Uinteger16{},
			wantErr: true,
		},
		{
			name:    "test uint16 slice pop out of bounds -1",
			args:    args{index: -1},
			i:       Uinteger16{1, 2, 3, 4, 5},
			want:    0,
			want1:   Uinteger16{},
			wantErr: true,
		},
		{
			name:    "test uint16 slice pop out of bounds 10",
			args:    args{index: 10},
			i:       Uinteger16{1, 2, 3, 4, 5},
			want:    0,
			want1:   Uinteger16{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.i.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uinteger16.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uinteger16.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Uinteger16.Pop() got1 = %v, want %v", got1, tt.want1)
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
			i:    []uint16{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "test empty true",
			i:    []uint16{},
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
