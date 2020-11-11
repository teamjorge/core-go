package slices

import (
	"reflect"
	"testing"
)

func TestConvert{{ .Slice.SliceName }}(t *testing.T) {
	{{ .Slice.SliceModifier }} := []{{ .Slice.SliceType }}{ {{ .TestItems }} }
	{{ .Slice.SliceModifier }} = {{ .Slice.SliceName }}({{ .Slice.SliceModifier }})
}

func Test{{ .Slice.SliceName }}_ForEach(t *testing.T) {
	testCache{{ .Slice.SliceName }}s := []{{ .Slice.SliceType }}{}

	type args struct {
		modifier func(index int, val {{ .Slice.SliceType }})
	}
	tests := []struct {
		name string
		{{ .Slice.SliceModifier }}          {{ .Slice.SliceName }}
		args       args
		cacheItems []{{ .Slice.SliceType }}
	}{
		{
			name: "test {{ .Slice.SliceType }} slice for each",
			args: args{
				modifier: func(index int, val {{ .Slice.SliceType }}) {
					testCache{{ .Slice.SliceName }}s = append(testCache{{ .Slice.SliceName }}s, val)
				},
			},
			{{ .Slice.SliceModifier }}:          {{ .Slice.SliceName }}{ {{ .TestItems }} },
			cacheItems: []{{ .Slice.SliceType }}{ {{ .TestItems }} },
		},
		{
			name: "test empty {{ .Slice.SliceType }} slice for each",
			args: args{
				modifier: func(index int, val {{ .Slice.SliceType }}) {
					testCache{{ .Slice.SliceName }}s = append(testCache{{ .Slice.SliceName }}s, val)
				},
			},
			{{ .Slice.SliceModifier }}:          {{ .Slice.SliceName }}{},
			cacheItems: []{{ .Slice.SliceType }}{},
		},
	}
	for _, tt := range tests {
		testCache{{ .Slice.SliceName }}s = []{{ .Slice.SliceType }}{}
		t.Run(tt.name, func(t *testing.T) {
			tt.{{ .Slice.SliceModifier }}.ForEach(tt.args.modifier)
		})
		if len(testCache{{ .Slice.SliceName }}s) != len(tt.cacheItems) {
			t.Errorf(
				"ForEach expected a cache length of %d but contains %d elements",
				len(tt.cacheItems),
				len(testCache{{ .Slice.SliceName }}s),
			)
		}

		if !reflect.DeepEqual(testCache{{ .Slice.SliceName }}s, tt.cacheItems) {
			t.Errorf("ForEach() = %v, want %v", testCache{{ .Slice.SliceName }}s, tt.cacheItems)
		}
	}
}

func Test{{ .Slice.SliceName }}_Map(t *testing.T) {
	type args struct {
		modifier func(index int, val {{ .Slice.SliceType }}) {{ .Slice.SliceType }}
	}
	tests := []struct {
		name string
		{{ .Slice.SliceModifier }}    {{ .Slice.SliceName }}
		args args
		want {{ .Slice.SliceName }}
	}{
		// TODO: Add Success Map Test
		{
			name: "test empty {{ .Slice.SliceType }} slice map",
			args: args{
				modifier: func(index int, val {{ .Slice.SliceType }}) {{ .Slice.SliceType }} {
					return {{ .NilValue }}
				},
			},
			{{ .Slice.SliceModifier }}:    {{ .Slice.SliceName }}{},
			want: []{{ .Slice.SliceType }}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.{{ .Slice.SliceModifier }}.Map(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("{{ .Slice.SliceName }}.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{ .Slice.SliceName }}_Filter(t *testing.T) {
	type args struct {
		modifier func(index int, val {{ .Slice.SliceType }}) bool
	}
	tests := []struct {
		name string
		{{ .Slice.SliceModifier }}    {{ .Slice.SliceName }}
		args args
		want {{ .Slice.SliceName }}
	}{
		{
			name: "test {{ .Slice.SliceType }} slice filter",
			args: args{
				modifier: func(index int, val {{ .Slice.SliceType }}) bool {
					return (val == {{ index .TestItemsSplit 1 }})
				},
			},
			{{ .Slice.SliceModifier }}:    {{ .Slice.SliceName }}{ {{ .TestItems }} },
			want: {{ .Slice.SliceName }}{ {{ index .TestItemsSplit 1 }} },
		},
		{
			name: "test empty {{ .Slice.SliceType }} slice filter",
			args: args{
				modifier: func(index int, val {{ .Slice.SliceType }}) bool {
					return true
				},
			},
			{{ .Slice.SliceModifier }}:    {{ .Slice.SliceName }}{},
			want: []{{ .Slice.SliceType }}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.{{ .Slice.SliceModifier }}.Filter(tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("{{ .Slice.SliceName }}.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{ .Slice.SliceName }}_Pop(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		{{ .Slice.SliceModifier }}       {{ .Slice.SliceName }}
		args    args
		want    {{ .Slice.SliceType }}
		want1   {{ .Slice.SliceName }}
		wantErr bool
	}{
		// TODO: Add test for popping first item
		// TODO: Add test for popping last item
		{
			name:    "test {{ .Slice.SliceType }} slice pop empty",
			args:    args{index: 0},
			{{ .Slice.SliceModifier }}:       {{ .Slice.SliceName }}{},
			want:    {{ .NilValue }},
			want1:   {{ .Slice.SliceName }}{},
			wantErr: true,
		},
		{
			name:    "test {{ .Slice.SliceType }} slice pop out of bounds -1",
			args:    args{index: -1},
			{{ .Slice.SliceModifier }}:       {{ .Slice.SliceName }}{ {{ .TestItems }} },
			want:    {{ .NilValue }},
			want1:   {{ .Slice.SliceName }}{},
			wantErr: true,
		},
		{
			name:    "test {{ .Slice.SliceType }} slice pop out of bounds 10",
			args:    args{index: 10},
			{{ .Slice.SliceModifier }}:       {{ .Slice.SliceName }}{ {{ .TestItems }} },
			want:    {{ .NilValue }},
			want1:   {{ .Slice.SliceName }}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.{{ .Slice.SliceModifier }}.Pop(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("{{ .Slice.SliceName }}.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("{{ .Slice.SliceName }}.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("{{ .Slice.SliceName }}.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test{{ .Slice.SliceName }}_Empty(t *testing.T) {
	tests := []struct {
		name string
		{{ .Slice.SliceModifier }}    {{ .Slice.SliceName }}
		want bool
	}{
		{
			name: "test empty false",
			{{ .Slice.SliceModifier }}:    []{{ .Slice.SliceType }}{ {{ .TestItems }} },
			want: false,
		},
		{
			name: "test empty true",
			{{ .Slice.SliceModifier }}:    []{{ .Slice.SliceType }}{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.{{ .Slice.SliceModifier }}.Empty(); got != tt.want {
				t.Errorf("{{ .Slice.SliceName }}.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
