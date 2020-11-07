package slices

import (
	"reflect"
	"testing"
)

func TestConvertBool(t *testing.T) {
	b := []bool{false, true}
	b = BoolSlice(b)
}

func TestBoolSlice_ForEach(t *testing.T) {
	testCacheBools := []bool{}

	type args struct {
		modifier func(index int, val bool)
	}
	tests := []struct {
		name       string
		b          BoolSlice
		args       args
		cacheItems []bool
	}{
		{
			name: "test bool slice for each",
			args: args{
				modifier: func(i int, b bool) {
					testCacheBools = append(testCacheBools, b)
				},
			},
			b:          BoolSlice{true, false, true, false},
			cacheItems: []bool{true, false, true, false},
		},
		{
			name: "test empty bool slice for each",
			args: args{
				modifier: func(i int, b bool) {
					testCacheBools = append(testCacheBools, b)
				},
			},
			b:          BoolSlice{},
			cacheItems: []bool{},
		},
		{
			name: "test bool slice for each invert",
			args: args{
				modifier: func(i int, b bool) {
					testCacheBools = append(testCacheBools, !b)
				},
			},
			b:          BoolSlice{true, false, true, false},
			cacheItems: []bool{false, true, false, true},
		},
	}
	for _, tt := range tests {
		testCacheBools = []bool{}
		t.Run(tt.name, func(t *testing.T) {
			tt.b.ForEach(tt.args.modifier)
		})
		if len(testCacheBools) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCacheBools),
			)
		}

		if !reflect.DeepEqual(testCacheBools, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCacheBools, tt.cacheItems)
		}
	}
}

func TestBoolSlice_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val bool) bool
	}
	tests := []struct {
		name string
		b    BoolSlice
		args args
		want BoolSlice
	}{
		{
			name: "test bool slice map",
			args: args{
				modifier: func(i int, b bool) bool {
					return b
				},
			},
			b:    BoolSlice{true, false, true, false},
			want: BoolSlice{true, false, true, false},
		},
		{
			name: "test empty bool slice map",
			args: args{
				modifier: func(i int, b bool) bool {
					return b
				},
			},
			b:    BoolSlice{},
			want: BoolSlice{},
		},
		{
			name: "test bool slice map invert",
			args: args{
				modifier: func(i int, b bool) bool {
					return !b
				},
			},
			b:    BoolSlice{true, false, true, false},
			want: BoolSlice{false, true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolSlice.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolSlice_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val bool) bool
	}
	tests := []struct {
		name string
		b    BoolSlice
		args args
		want BoolSlice
	}{
		{
			name: "test bool slice filter",
			args: args{
				modifier: func(i int, b bool) bool {
					return b
				},
			},
			b:    BoolSlice{true, false, true, false},
			want: BoolSlice{true, true},
		},
		{
			name: "test empty bool slice filter",
			args: args{
				modifier: func(i int, b bool) bool {
					return b
				},
			},
			b:    BoolSlice{},
			want: BoolSlice{},
		},
		{
			name: "test bool slice map invert",
			args: args{
				modifier: func(i int, b bool) bool {
					return !b
				},
			},
			b:    BoolSlice{true, false, true, false},
			want: BoolSlice{false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolSlice.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolSlice_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		b       BoolSlice
		args    args
		want    bool
		want1   BoolSlice
		wantErr bool
	}{
		{
			name:    "test bool slice pop first",
			args:    args{index: 0},
			b:       BoolSlice{true, false, false, false, true, true, false, true},
			want:    true,
			want1:   BoolSlice{false, false, false, true, true, false, true},
			wantErr: false,
		},
		{
			name:    "test bool slice pop last",
			args:    args{index: 8},
			b:       BoolSlice{true, false, true, false, true, false, true, false, false},
			want:    false,
			want1:   BoolSlice{true, false, true, false, true, false, true, false},
			wantErr: false,
		},
		{
			name:    "test bool slice pop empty",
			args:    args{index: 0},
			b:       BoolSlice{},
			want:    false,
			want1:   BoolSlice{},
			wantErr: true,
		},
		{
			name:    "test bool slice pop out of bounds -1",
			args:    args{index: -1},
			b:       BoolSlice{false, false, false},
			want:    false,
			want1:   BoolSlice{},
			wantErr: true,
		},
		{
			name:    "test bool slice pop out of bounds 10",
			args:    args{index: 10},
			b:       BoolSlice{false, false, false},
			want:    false,
			want1:   BoolSlice{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.b.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("BoolSlice.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BoolSlice.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BoolSlice.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAny(t *testing.T) {
	type args struct {
		in BoolSlice
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test bool slice any true",
			args: args{in: BoolSlice{false, false, false, true, false}},
			want: true,
		},
		{
			name: "test bool slice any false",
			args: args{in: BoolSlice{false, false, false, false, false, false}},
			want: false,
		},
		{
			name: "test bool slice any empty",
			args: args{in: BoolSlice{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.args.in); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	type args struct {
		in BoolSlice
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test bool slice all true",
			args: args{in: BoolSlice{true, true, true, true, true}},
			want: true,
		},
		{
			name: "test bool slice all false",
			args: args{in: BoolSlice{true, true, false, true, true}},
			want: false,
		},
		{
			name: "test bool slice all empty",
			args: args{in: BoolSlice{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.args.in); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolSlice_Empty(t *testing.T) {
	tests := []struct {
		name string
		b    BoolSlice
		want bool
	}{
		{
			name: "test empty false",
			b:    []bool{true, false},
			want: false,
		},
		{
			name: "test empty true",
			b:    []bool{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Empty(); got != tt.want {
				t.Errorf("BoolSlice.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}