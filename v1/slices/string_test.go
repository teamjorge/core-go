package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertString(t *testing.T) {
	s := []string{"ha", "what"}
	s = StringSlice(s)
}

func TestStringSlice_ForEach(t *testing.T) {
	testCacheStrings := []string{}

	type args struct {
		modifier func(index int, s string)
	}
	tests := []struct {
		name       string
		s          StringSlice
		args       args
		cacheItems []string
	}{
		{
			name: "test string slice for each",
			args: args{
				modifier: func(i int, s string) {
					testCacheStrings = append(testCacheStrings, s)
				},
			},
			s:          StringSlice{"this", "is"},
			cacheItems: []string{"this", "is"},
		},
		{
			name: "test empty string slice for each",
			args: args{
				modifier: func(i int, s string) {
					testCacheStrings = append(testCacheStrings, s)
				},
			},
			s:          StringSlice{},
			cacheItems: []string{},
		},
		{
			name: "test modify string slice for each",
			args: args{
				modifier: func(i int, s string) {
					s += "_lol"
					testCacheStrings = append(testCacheStrings, s)
				},
			},
			s:          StringSlice{"this", "is"},
			cacheItems: []string{"this_lol", "is_lol"},
		},
	}
	for _, tt := range tests {
		testCacheStrings = []string{}
		t.Run(tt.name, func(t *testing.T) {
			tt.s.ForEach(tt.args.modifier)
		})
		if len(testCacheStrings) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCacheStrings),
			)
		}

		if !reflect.DeepEqual(testCacheStrings, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCacheStrings, tt.cacheItems)
		}
	}
}

func TestStringSlice_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val string) string
	}
	tests := []struct {
		name string
		s    StringSlice
		args args
		want StringSlice
	}{
		{
			name: "test string slice map",
			args: args{
				modifier: func(i int, s string) string {
					return fmt.Sprintf("%s_%d", s, i)
				},
			},
			s:    StringSlice{"this", "is"},
			want: []string{"this_0", "is_1"},
		},
		{
			name: "test empty string slice map",
			args: args{
				modifier: func(i int, s string) string {
					return fmt.Sprintf("%s_%d", s, i)
				},
			},
			s:    StringSlice{},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSlice.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSlice_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val string) bool
	}
	tests := []struct {
		name string
		s    StringSlice
		args args
		want StringSlice
	}{
		{
			name: "test string slice filter",
			args: args{
				modifier: func(i int, s string) bool {
					return !(s == "this")
				},
			},
			s:    StringSlice{"this", "is"},
			want: StringSlice{"is"},
		},
		{
			name: "test empty string slice filter",
			args: args{
				modifier: func(i int, s string) bool {
					return true
				},
			},
			s:    StringSlice{},
			want: []string{},
		},
		{
			name: "test string slice filter for empty",
			args: args{
				modifier: func(i int, s string) bool {
					return !(s == "")
				},
			},
			s:    StringSlice{"this", "", "is", "", "", "", "test"},
			want: StringSlice{"this", "is", "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSlice.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSlice_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		s       StringSlice
		args    args
		want    string
		want1   StringSlice
		wantErr bool
	}{
		{
			name:    "test string slice pop first",
			args:    args{index: 0},
			s:       StringSlice{"this", "", "is", "", "", "", "test"},
			want:    "this",
			want1:   StringSlice{"", "is", "", "", "", "test"},
			wantErr: false,
		},
		{
			name:    "test string slice pop last",
			args:    args{index: 6},
			s:       StringSlice{"this", "", "is", "", "", "", "test"},
			want:    "test",
			want1:   StringSlice{"this", "", "is", "", "", ""},
			wantErr: false,
		},
		{
			name:    "test string slice pop empty",
			args:    args{index: 0},
			s:       StringSlice{},
			want:    "",
			want1:   StringSlice{},
			wantErr: true,
		},
		{
			name:    "test string slice pop out of bounds -1",
			args:    args{index: -1},
			s:       StringSlice{"this", "", "is", "", "", "", "test"},
			want:    "",
			want1:   StringSlice{},
			wantErr: true,
		},
		{
			name:    "test string slice pop out of bounds 10",
			args:    args{index: 10},
			s:       StringSlice{"this", "", "is", "", "", "", "test"},
			want:    "",
			want1:   StringSlice{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringSlice.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringSlice.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StringSlice.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStringSlice_Empty(t *testing.T) {
	tests := []struct {
		name string
		s    StringSlice
		want bool
	}{
		{
			name: "test empty false",
			s:    []string{"this", "is"},
			want: false,
		},
		{
			name: "test empty true",
			s:    []string{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Empty(); got != tt.want {
				t.Errorf("StringSlice.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
