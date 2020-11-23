package slices

import (
	"reflect"
	"testing"
)

func TestConvertInteger(t *testing.T) {
	i := []int{1, 2, 3, 4, 5}
	i = Integer(i)
}

func TestInteger_ForEach(t *testing.T) {
	testCacheIntegers := []int{}

	type args struct {
		modifier func(index int, val int)
	}
	tests := []struct {
		name       string
		i          Integer
		args       args
		cacheItems []int
	}{
		{
			name: "test int slice for each",
			args: args{
				modifier: func(index int, val int) {
					testCacheIntegers = append(testCacheIntegers, val)
				},
			},
			i:          Integer{1, 2, 3, 4, 5},
			cacheItems: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test empty int slice for each",
			args: args{
				modifier: func(index int, val int) {
					testCacheIntegers = append(testCacheIntegers, val)
				},
			},
			i:          Integer{},
			cacheItems: []int{},
		},
	}
	for _, tt := range tests {
		testCacheIntegers = []int{}
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ForEach(tt.args.modifier)
		})
		if len(testCacheIntegers) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCacheIntegers),
			)
		}

		if !reflect.DeepEqual(testCacheIntegers, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCacheIntegers, tt.cacheItems)
		}
	}
}

func TestInteger_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val int) int
	}
	tests := []struct {
		name string
		i    Integer
		args args
		want Integer
	}{
		{
			name: "test success int slice map",
			args: args{
				modifier: func(index int, val int) int {
					return val + 1
				},
			},
			i:    Integer{1, 2, 3, 4, 5},
			want: []int{2, 3, 4, 5, 6},
		},
		{
			name: "test empty int slice map",
			args: args{
				modifier: func(index int, val int) int {
					return 0
				},
			},
			i:    Integer{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integer.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val int) bool
	}
	tests := []struct {
		name string
		i    Integer
		args args
		want Integer
	}{
		{
			name: "test int slice filter",
			args: args{
				modifier: func(index int, val int) bool {
					return (val == 2)
				},
			},
			i:    Integer{1, 2, 3, 4, 5},
			want: Integer{2},
		},
		{
			name: "test empty int slice filter",
			args: args{
				modifier: func(index int, val int) bool {
					return true
				},
			},
			i:    Integer{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integer.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		i       Integer
		args    args
		want    int
		want1   Integer
		wantErr bool
	}{
		{
			name:    "test uint slice pop first",
			args:    args{index: 0},
			i:       Integer{1, 2, 3, 4, 5},
			want:    1,
			want1:   Integer{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "test uint slice pop last",
			args:    args{index: 4},
			i:       Integer{1, 2, 3, 4, 5},
			want:    5,
			want1:   Integer{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "test int slice pop empty",
			args:    args{index: 0},
			i:       Integer{},
			want:    0,
			want1:   Integer{},
			wantErr: true,
		},
		{
			name:    "test int slice pop out of bounds -1",
			args:    args{index: -1},
			i:       Integer{1, 2, 3, 4, 5},
			want:    0,
			want1:   Integer{},
			wantErr: true,
		},
		{
			name:    "test int slice pop out of bounds 10",
			args:    args{index: 10},
			i:       Integer{1, 2, 3, 4, 5},
			want:    0,
			want1:   Integer{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.i.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Integer.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Integer.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Integer.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInteger_Empty(t *testing.T) {
	tests := []struct {
		name string
		i    Integer
		want bool
	}{
		{
			name: "test empty false",
			i:    []int{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "test empty true",
			i:    []int{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Empty(); got != tt.want {
				t.Errorf("Integer.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
