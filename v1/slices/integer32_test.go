package slices

import (
	"reflect"
	"testing"
)

func TestConvertInteger32(t *testing.T) {
	i := []int32{1, 2, 3, 4, 5}
	i = Integer32(i)
}

func TestInteger32_ForEach(t *testing.T) {
	testCacheInteger32s := []int32{}

	type args struct {
		modifier func(index int, val int32)
	}
	tests := []struct {
		name       string
		i          Integer32
		args       args
		cacheItems []int32
	}{
		{
			name: "test int32 slice for each",
			args: args{
				modifier: func(index int, val int32) {
					testCacheInteger32s = append(testCacheInteger32s, val)
				},
			},
			i:          Integer32{1, 2, 3, 4, 5},
			cacheItems: []int32{1, 2, 3, 4, 5},
		},
		{
			name: "test empty int32 slice for each",
			args: args{
				modifier: func(index int, val int32) {
					testCacheInteger32s = append(testCacheInteger32s, val)
				},
			},
			i:          Integer32{},
			cacheItems: []int32{},
		},
	}
	for _, tt := range tests {
		testCacheInteger32s = []int32{}
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ForEach(tt.args.modifier)
		})
		if len(testCacheInteger32s) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCacheInteger32s),
			)
		}

		if !reflect.DeepEqual(testCacheInteger32s, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCacheInteger32s, tt.cacheItems)
		}
	}
}

func TestInteger32_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val int32) int32
	}
	tests := []struct {
		name string
		i    Integer32
		args args
		want Integer32
	}{
		{
			name: "test success int32 slice map",
			args: args{
				modifier: func(index int, val int32) int32 {
					return val + 1
				},
			},
			i:    Integer32{1, 2, 3, 4, 5},
			want: []int32{2, 3, 4, 5, 6},
		},
		{
			name: "test empty int32 slice map",
			args: args{
				modifier: func(index int, val int32) int32 {
					return 0
				},
			},
			i:    Integer32{},
			want: []int32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integer32.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger32_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val int32) bool
	}
	tests := []struct {
		name string
		i    Integer32
		args args
		want Integer32
	}{
		{
			name: "test int32 slice filter",
			args: args{
				modifier: func(index int, val int32) bool {
					return (val == 2)
				},
			},
			i:    Integer32{1, 2, 3, 4, 5},
			want: Integer32{2},
		},
		{
			name: "test empty int32 slice filter",
			args: args{
				modifier: func(index int, val int32) bool {
					return true
				},
			},
			i:    Integer32{},
			want: []int32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integer32.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger32_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		i       Integer32
		args    args
		want    int32
		want1   Integer32
		wantErr bool
	}{
		{
			name:    "test uint32 slice pop first",
			args:    args{index: 0},
			i:       Integer32{1, 2, 3, 4, 5},
			want:    1,
			want1:   Integer32{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "test uint32 slice pop last",
			args:    args{index: 4},
			i:       Integer32{1, 2, 3, 4, 5},
			want:    5,
			want1:   Integer32{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "test int32 slice pop empty",
			args:    args{index: 0},
			i:       Integer32{},
			want:    0,
			want1:   Integer32{},
			wantErr: true,
		},
		{
			name:    "test int32 slice pop out of bounds -1",
			args:    args{index: -1},
			i:       Integer32{1, 2, 3, 4, 5},
			want:    0,
			want1:   Integer32{},
			wantErr: true,
		},
		{
			name:    "test int32 slice pop out of bounds 10",
			args:    args{index: 10},
			i:       Integer32{1, 2, 3, 4, 5},
			want:    0,
			want1:   Integer32{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.i.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Integer32.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Integer32.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Integer32.Pop() got1 = %v, want %v", got1, tt.want1)
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
			i:    []int32{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "test empty true",
			i:    []int32{},
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

func TestInteger32_Contains(t *testing.T) {
	type args struct {
		value int32
	}
	tests := []struct {
		name string
		i    Integer32
		args args
		want bool
	}{
		{
			name: "test does contain",
			i:    Integer32{1, 2, 3, 4, 5},
			args: args{value: 3},
			want: true,
		},
		{
			name: "test does not contain",
			i:    Integer32{1, 2, 3, 4, 5},
			args: args{value: 12},
			want: false,
		},
		{
			name: "test contain empty",
			i:    Integer32{},
			args: args{value: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Contains(tt.args.value); got != tt.want {
				t.Errorf("Integer32.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
