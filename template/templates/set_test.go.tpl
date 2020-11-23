package sets

import (
	"reflect"
	"sort"
	"testing"
)

func TestTo{{ .Set.SetName }}Slice(t *testing.T) {
	type args struct {
		in [][]{{ .Set.SetType }}
	}
	tests := []struct {
		name string
		args args
		want []{{ .Set.SetType }}
	}{
		{
			name: "test single slice",
			args: args{in: [][]{{ .Set.SetType }}{ { {{ .TestItems }} } }},
			want: []{{ .Set.SetType }}{ {{ .TestItems }} },
		},
		{
			name: "test multiple slice",
			args: args{in: [][]{{ .Set.SetType }}{
				{ {{ .TestItems }} },
				{ {{ .TestItems }} },
			}},
			want: []{{ .Set.SetType }}{ {{ .TestItems }} },
		},
		{
			name: "test empty array",
			args: args{in: [][]{{ .Set.SetType }}{}},
			want: []{{ .Set.SetType }}{},
		},
		{
			name: "test empty args",
			args: args{},
			want: []{{ .Set.SetType }}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := To{{ .Set.SetName }}Slice(tt.args.in...)
			if len(got) != len(tt.want) {
				t.Errorf("Expected a result of length %d but received %d", len(tt.want), len(got))
			}
            sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("{{ .Set.SetName }}() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew{{ .Set.SetName }}(t *testing.T) {
	type args struct {
		in [][]{{ .Set.SetType }}
	}
	tests := []struct {
		name string
		args args
		want *{{ .Set.SetName }}
	}{
		{
			name: "test create new {{ .Set.SetType }} set simple",
			args: args{
				in: [][]{{ .Set.SetType }}{ { {{ .TestItems }} } },
			},
			want: &{{ .Set.SetName }}{ {{ .TestItemsMap }} },
		},
		{
			name: "test create new {{ .Set.SetType }} set empty",
			args: args{
				in: [][]{{ .Set.SetType }}{ { } },
			},
			want: &{{ .Set.SetName }}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New{{ .Set.SetName }}(tt.args.in...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New{{ .Set.SetName }}() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{ .Set.SetName }}_Add(t *testing.T) {
	type args struct {
		in []{{ .Set.SetType }}
	}
	tests := []struct {
		name string
		{{ .Set.SetModifier }}    {{ .Set.SetName }}
		args args
		want {{ .Set.SetName }}
	}{
		{
			name: "test add simple",
			args: args{
				in: []{{ .Set.SetType }}{ {{ .NilValue }} },
			},
			{{ .Set.SetModifier }}: {{ .Set.SetName }}{ {{ .TestItemsMap }} },
			want: {{ .Set.SetName }}{ {{ .TestItemsMapWithNil }} },
		},
		{
			name: "test add empty",
			args: args{
				in: []{{ .Set.SetType }}{},
			},
			{{ .Set.SetModifier }}: {{ .Set.SetName }}{ {{ .TestItemsMap }} },
			want: {{ .Set.SetName }}{ {{ .TestItemsMap }} },
		},
		{
			name: "test add same",
			args: args{
				in: []{{ .Set.SetType }}{ {{ .TestItems }} },
			},
			{{ .Set.SetModifier }}: {{ .Set.SetName }}{ {{ .TestItemsMap }} },
			want: {{ .Set.SetName }}{ {{ .TestItemsMap }} },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.{{ .Set.SetModifier }}.Add(tt.args.in...)
			if !reflect.DeepEqual(tt.{{ .Set.SetModifier }}, tt.want) {
				t.Errorf("Add() = %v, want %v", tt.{{ .Set.SetModifier }}, tt.want)
			}
		})
	}
}

func Test{{ .Set.SetName }}_Remove(t *testing.T) {
	type args struct {
		elem []{{ .Set.SetType }}
	}
	tests := []struct {
		name string
		{{ .Set.SetModifier }}    {{ .Set.SetName }}
		args args
		want {{ .Set.SetName }}
	}{
		{
			name: "test remove simple",
			args: args{
				elem: []{{ .Set.SetType }}{ {{ .NilValue }} },
			},
			{{ .Set.SetModifier }}: {{ .Set.SetName }}{ {{ .TestItemsMapWithNil }} },
			want: {{ .Set.SetName }}{ {{ .TestItemsMap }} },
		},
		{
			name: "test remove empty",
			args: args{
				elem: []{{ .Set.SetType }}{ {{ .NilValue }} },
			},
			{{ .Set.SetModifier }}: {{ .Set.SetName }}{},
			want: {{ .Set.SetName }}{},
		},
		{
			name: "test remove nothing",
			args: args{
				elem: []{{ .Set.SetType }}{},
			},
			{{ .Set.SetModifier }}: {{ .Set.SetName }}{ {{ .TestItemsMap }} },
			want: {{ .Set.SetName }}{ {{ .TestItemsMap }} },
		},
		{
			name: "test remove not exists",
			args: args{
				elem: []{{ .Set.SetType }}{ {{ .NilValue }} },
			},
			{{ .Set.SetModifier }}: {{ .Set.SetName }}{ {{ .TestItemsMap }} },
			want: {{ .Set.SetName }}{ {{ .TestItemsMap }} },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.{{ .Set.SetModifier }}.Remove(tt.args.elem...)
			if !reflect.DeepEqual(tt.{{ .Set.SetModifier }}, tt.want) {
				t.Errorf("Remove() = %v, want %v", tt.{{ .Set.SetModifier }}, tt.want)
			}
		})
	}
}

func Test{{ .Set.SetName }}_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		{{ .Set.SetModifier }}    {{ .Set.SetName }}
		want []{{ .Set.SetType }}
	}{
		{
			name: "test to slice simple",
			{{ .Set.SetModifier }}:    {{ .Set.SetName }}{ {{ .TestItemsMap }} },
			want: []{{ .Set.SetType }}{ {{ .TestItems }} },
		},
		{
			name: "test to slice empty",
			{{ .Set.SetModifier }}:    {{ .Set.SetName }}{},
			want: []{{ .Set.SetType }}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.{{ .Set.SetModifier }}.ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Expected a result of length %d but received %d", len(tt.want), len(got))
			}
            // If sort doesn't work, add custom sorting
            sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{ .Set.SetName }}_Empty(t *testing.T) {
	tests := []struct {
		name string
		{{ .Set.SetModifier }}    {{ .Set.SetName }}
		want bool
	}{
		{
			name: "test empty false",
			{{ .Set.SetModifier }}:    {{ .Set.SetName }}{ {{ .TestItemsMap }} },
			want: false,
		},
		{
			name: "test empty true",
			{{ .Set.SetModifier }}:    {{ .Set.SetName }}{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.{{ .Set.SetModifier }}.Empty(); got != tt.want {
				t.Errorf("{{ .Set.SetName }}.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
