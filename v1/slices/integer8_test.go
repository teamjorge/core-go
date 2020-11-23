package slices

import (
	"reflect"
	"testing"
)

func TestConvertInteger8(t *testing.T) {
	i := []int8{1, 2, 3, 4, 5}
	i = Integer8(i)
}

func TestInteger8_ForEach(t *testing.T) {
	testCacheInteger8s := []int8{}

	type args struct {
		modifier func(index int, val int8)
	}
	tests := []struct {
		name       string
		i          Integer8
		args       args
		cacheItems []int8
	}{
		{
			name: "test int8 slice for each",
			args: args{
				modifier: func(index int, val int8) {
					testCacheInteger8s = append(testCacheInteger8s, val)
				},
			},
			i:          Integer8{1, 2, 3, 4, 5},
			cacheItems: []int8{1, 2, 3, 4, 5},
		},
		{
			name: "test empty int8 slice for each",
			args: args{
				modifier: func(index int, val int8) {
					testCacheInteger8s = append(testCacheInteger8s, val)
				},
			},
			i:          Integer8{},
			cacheItems: []int8{},
		},
	}
	for _, tt := range tests {
		testCacheInteger8s = []int8{}
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ForEach(tt.args.modifier)
		})
		if len(testCacheInteger8s) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCacheInteger8s),
			)
		}

		if !reflect.DeepEqual(testCacheInteger8s, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCacheInteger8s, tt.cacheItems)
		}
	}
}

func TestInteger8_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val int8) int8
	}
	tests := []struct {
		name string
		i    Integer8
		args args
		want Integer8
	}{
		{
			name: "test success int8 slice map",
			args: args{
				modifier: func(index int, val int8) int8 {
					return val + 1
				},
			},
			i:    Integer8{1, 2, 3, 4, 5},
			want: []int8{2, 3, 4, 5, 6},
		},
		{
			name: "test empty int8 slice map",
			args: args{
				modifier: func(index int, val int8) int8 {
					return 0
				},
			},
			i:    Integer8{},
			want: []int8{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integer8.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger8_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val int8) bool
	}
	tests := []struct {
		name string
		i    Integer8
		args args
		want Integer8
	}{
		{
			name: "test int8 slice filter",
			args: args{
				modifier: func(index int, val int8) bool {
					return (val == 2)
				},
			},
			i:    Integer8{1, 2, 3, 4, 5},
			want: Integer8{2},
		},
		{
			name: "test empty int8 slice filter",
			args: args{
				modifier: func(index int, val int8) bool {
					return true
				},
			},
			i:    Integer8{},
			want: []int8{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integer8.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger8_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		i       Integer8
		args    args
		want    int8
		want1   Integer8
		wantErr bool
	}{
		{
			name:    "test uint8 slice pop first",
			args:    args{index: 0},
			i:       Integer8{1, 2, 3, 4, 5},
			want:    1,
			want1:   Integer8{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "test uint8 slice pop last",
			args:    args{index: 4},
			i:       Integer8{1, 2, 3, 4, 5},
			want:    5,
			want1:   Integer8{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "test int8 slice pop empty",
			args:    args{index: 0},
			i:       Integer8{},
			want:    0,
			want1:   Integer8{},
			wantErr: true,
		},
		{
			name:    "test int8 slice pop out of bounds -1",
			args:    args{index: -1},
			i:       Integer8{1, 2, 3, 4, 5},
			want:    0,
			want1:   Integer8{},
			wantErr: true,
		},
		{
			name:    "test int8 slice pop out of bounds 10",
			args:    args{index: 10},
			i:       Integer8{1, 2, 3, 4, 5},
			want:    0,
			want1:   Integer8{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.i.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Integer8.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Integer8.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Integer8.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInteger8_Empty(t *testing.T) {
	tests := []struct {
		name string
		i    Integer8
		want bool
	}{
		{
			name: "test empty false",
			i:    []int8{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "test empty true",
			i:    []int8{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Empty(); got != tt.want {
				t.Errorf("Integer8.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
