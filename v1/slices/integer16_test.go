package slices

import (
	"reflect"
	"testing"
)

func TestConvertInteger16(t *testing.T) {
	i := []int16{1, 2, 3, 4, 5}
	i = Integer16(i)
}

func TestInteger16_ForEach(t *testing.T) {
	testCacheInteger16s := []int16{}

	type args struct {
		modifier func(index int, val int16)
	}
	tests := []struct {
		name       string
		i          Integer16
		args       args
		cacheItems []int16
	}{
		{
			name: "test int16 slice for each",
			args: args{
				modifier: func(index int, val int16) {
					testCacheInteger16s = append(testCacheInteger16s, val)
				},
			},
			i:          Integer16{1, 2, 3, 4, 5},
			cacheItems: []int16{1, 2, 3, 4, 5},
		},
		{
			name: "test empty int16 slice for each",
			args: args{
				modifier: func(index int, val int16) {
					testCacheInteger16s = append(testCacheInteger16s, val)
				},
			},
			i:          Integer16{},
			cacheItems: []int16{},
		},
	}
	for _, tt := range tests {
		testCacheInteger16s = []int16{}
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ForEach(tt.args.modifier)
		})
		if len(testCacheInteger16s) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCacheInteger16s),
			)
		}

		if !reflect.DeepEqual(testCacheInteger16s, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCacheInteger16s, tt.cacheItems)
		}
	}
}

func TestInteger16_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val int16) int16
	}
	tests := []struct {
		name string
		i    Integer16
		args args
		want Integer16
	}{
		{
			name: "test success int16 slice map",
			args: args{
				modifier: func(index int, val int16) int16 {
					return val + 1
				},
			},
			i:    Integer16{1, 2, 3, 4, 5},
			want: []int16{2, 3, 4, 5, 6},
		},
		{
			name: "test empty int16 slice map",
			args: args{
				modifier: func(index int, val int16) int16 {
					return 0
				},
			},
			i:    Integer16{},
			want: []int16{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integer16.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger16_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val int16) bool
	}
	tests := []struct {
		name string
		i    Integer16
		args args
		want Integer16
	}{
		{
			name: "test int16 slice filter",
			args: args{
				modifier: func(index int, val int16) bool {
					return (val == 2)
				},
			},
			i:    Integer16{1, 2, 3, 4, 5},
			want: Integer16{2},
		},
		{
			name: "test empty int16 slice filter",
			args: args{
				modifier: func(index int, val int16) bool {
					return true
				},
			},
			i:    Integer16{},
			want: []int16{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integer16.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger16_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		i       Integer16
		args    args
		want    int16
		want1   Integer16
		wantErr bool
	}{
		{
			name:    "test uint16 slice pop first",
			args:    args{index: 0},
			i:       Integer16{1, 2, 3, 4, 5},
			want:    1,
			want1:   Integer16{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "test uint16 slice pop last",
			args:    args{index: 4},
			i:       Integer16{1, 2, 3, 4, 5},
			want:    5,
			want1:   Integer16{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "test int16 slice pop empty",
			args:    args{index: 0},
			i:       Integer16{},
			want:    0,
			want1:   Integer16{},
			wantErr: true,
		},
		{
			name:    "test int16 slice pop out of bounds -1",
			args:    args{index: -1},
			i:       Integer16{1, 2, 3, 4, 5},
			want:    0,
			want1:   Integer16{},
			wantErr: true,
		},
		{
			name:    "test int16 slice pop out of bounds 10",
			args:    args{index: 10},
			i:       Integer16{1, 2, 3, 4, 5},
			want:    0,
			want1:   Integer16{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.i.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Integer16.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Integer16.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Integer16.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInteger16_Empty(t *testing.T) {
	tests := []struct {
		name string
		i    Integer16
		want bool
	}{
		{
			name: "test empty false",
			i:    []int16{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "test empty true",
			i:    []int16{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Empty(); got != tt.want {
				t.Errorf("Integer16.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
