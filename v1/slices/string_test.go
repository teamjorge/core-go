package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertString(t *testing.T) {
	s := []string{"ha", "what"}
	s = String(s)
}

func TestString_ForEach(t *testing.T) {
	testCacheStrings := []string{}

	type args struct {
		modifier func(index int, s string)
	}
	tests := []struct {
		name       string
		s          String
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
			s:          String{"this", "is"},
			cacheItems: []string{"this", "is"},
		},
		{
			name: "test empty string slice for each",
			args: args{
				modifier: func(i int, s string) {
					testCacheStrings = append(testCacheStrings, s)
				},
			},
			s:          String{},
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
			s:          String{"this", "is"},
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

func TestString_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val string) string
	}
	tests := []struct {
		name string
		s    String
		args args
		want String
	}{
		{
			name: "test string slice map",
			args: args{
				modifier: func(i int, s string) string {
					return fmt.Sprintf("%s_%d", s, i)
				},
			},
			s:    String{"this", "is"},
			want: []string{"this_0", "is_1"},
		},
		{
			name: "test empty string slice map",
			args: args{
				modifier: func(i int, s string) string {
					return fmt.Sprintf("%s_%d", s, i)
				},
			},
			s:    String{},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val string) bool
	}
	tests := []struct {
		name string
		s    String
		args args
		want String
	}{
		{
			name: "test string slice filter",
			args: args{
				modifier: func(i int, s string) bool {
					return !(s == "this")
				},
			},
			s:    String{"this", "is"},
			want: String{"is"},
		},
		{
			name: "test empty string slice filter",
			args: args{
				modifier: func(i int, s string) bool {
					return true
				},
			},
			s:    String{},
			want: []string{},
		},
		{
			name: "test string slice filter for empty",
			args: args{
				modifier: func(i int, s string) bool {
					return !(s == "")
				},
			},
			s:    String{"this", "", "is", "", "", "", "test"},
			want: String{"this", "is", "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		s       String
		args    args
		want    string
		want1   String
		wantErr bool
	}{
		{
			name:    "test string slice pop first",
			args:    args{index: 0},
			s:       String{"this", "", "is", "", "", "", "test"},
			want:    "this",
			want1:   String{"", "is", "", "", "", "test"},
			wantErr: false,
		},
		{
			name:    "test string slice pop last",
			args:    args{index: 6},
			s:       String{"this", "", "is", "", "", "", "test"},
			want:    "test",
			want1:   String{"this", "", "is", "", "", ""},
			wantErr: false,
		},
		{
			name:    "test string slice pop empty",
			args:    args{index: 0},
			s:       String{},
			want:    "",
			want1:   String{},
			wantErr: true,
		},
		{
			name:    "test string slice pop out of bounds -1",
			args:    args{index: -1},
			s:       String{"this", "", "is", "", "", "", "test"},
			want:    "",
			want1:   String{},
			wantErr: true,
		},
		{
			name:    "test string slice pop out of bounds 10",
			args:    args{index: 10},
			s:       String{"this", "", "is", "", "", "", "test"},
			want:    "",
			want1:   String{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("String.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("String.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("String.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestString_Empty(t *testing.T) {
	tests := []struct {
		name string
		s    String
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
				t.Errorf("String.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
