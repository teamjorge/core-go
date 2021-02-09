package slices

import (
	"reflect"
	"testing"
)

func TestConvertInteger64(t *testing.T) {
	i := []int64{1, 2, 3, 4, 5}
	i = Integer64(i)
}

func TestInteger64_ForEach(t *testing.T) {
	testCacheInteger64s := []int64{}

	type args struct {
		modifier func(index int, val int64)
	}
	tests := []struct {
		name       string
		i          Integer64
		args       args
		cacheItems []int64
	}{
		{
			name: "test int64 slice for each",
			args: args{
				modifier: func(index int, val int64) {
					testCacheInteger64s = append(testCacheInteger64s, val)
				},
			},
			i:          Integer64{1, 2, 3, 4, 5},
			cacheItems: []int64{1, 2, 3, 4, 5},
		},
		{
			name: "test empty int64 slice for each",
			args: args{
				modifier: func(index int, val int64) {
					testCacheInteger64s = append(testCacheInteger64s, val)
				},
			},
			i:          Integer64{},
			cacheItems: []int64{},
		},
	}
	for _, tt := range tests {
		testCacheInteger64s = []int64{}
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ForEach(tt.args.modifier)
		})
		if len(testCacheInteger64s) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCacheInteger64s),
			)
		}

		if !reflect.DeepEqual(testCacheInteger64s, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCacheInteger64s, tt.cacheItems)
		}
	}
}

func TestInteger64_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val int64) int64
	}
	tests := []struct {
		name string
		i    Integer64
		args args
		want Integer64
	}{
		{
			name: "test success int64 slice map",
			args: args{
				modifier: func(index int, val int64) int64 {
					return val + 1
				},
			},
			i:    Integer64{1, 2, 3, 4, 5},
			want: []int64{2, 3, 4, 5, 6},
		},
		{
			name: "test empty int64 slice map",
			args: args{
				modifier: func(index int, val int64) int64 {
					return 0
				},
			},
			i:    Integer64{},
			want: []int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integer64.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val int64) bool
	}
	tests := []struct {
		name string
		i    Integer64
		args args
		want Integer64
	}{
		{
			name: "test int64 slice filter",
			args: args{
				modifier: func(index int, val int64) bool {
					return (val == 2)
				},
			},
			i:    Integer64{1, 2, 3, 4, 5},
			want: Integer64{2},
		},
		{
			name: "test empty int64 slice filter",
			args: args{
				modifier: func(index int, val int64) bool {
					return true
				},
			},
			i:    Integer64{},
			want: []int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integer64.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		i       Integer64
		args    args
		want    int64
		want1   Integer64
		wantErr bool
	}{
		{
			name:    "test uint64 slice pop first",
			args:    args{index: 0},
			i:       Integer64{1, 2, 3, 4, 5},
			want:    1,
			want1:   Integer64{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "test uint64 slice pop last",
			args:    args{index: 4},
			i:       Integer64{1, 2, 3, 4, 5},
			want:    5,
			want1:   Integer64{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "test int64 slice pop empty",
			args:    args{index: 0},
			i:       Integer64{},
			want:    0,
			want1:   Integer64{},
			wantErr: true,
		},
		{
			name:    "test int64 slice pop out of bounds -1",
			args:    args{index: -1},
			i:       Integer64{1, 2, 3, 4, 5},
			want:    0,
			want1:   Integer64{},
			wantErr: true,
		},
		{
			name:    "test int64 slice pop out of bounds 10",
			args:    args{index: 10},
			i:       Integer64{1, 2, 3, 4, 5},
			want:    0,
			want1:   Integer64{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.i.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Integer64.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Integer64.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Integer64.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInteger64_Empty(t *testing.T) {
	tests := []struct {
		name string
		i    Integer64
		want bool
	}{
		{
			name: "test empty false",
			i:    []int64{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "test empty true",
			i:    []int64{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Empty(); got != tt.want {
				t.Errorf("Integer64.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_Contains(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name string
		i    Integer64
		args args
		want bool
	}{
		{
			name: "test does contain",
			i:    Integer64{1, 2, 3, 4, 5},
			args: args{value: 3},
			want: true,
		},
		{
			name: "test does not contain",
			i:    Integer64{1, 2, 3, 4, 5},
			args: args{value: 12},
			want: false,
		},
		{
			name: "test contain empty",
			i:    Integer64{},
			args: args{value: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Contains(tt.args.value); got != tt.want {
				t.Errorf("Integer64.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
