package slices

import (
	"reflect"
	"testing"
)

// Person is a struct for testing generic methods in the slices package
type Person struct {
	Name string
}

// Persons is a slice of Person instances
type Persons []Person

func (m Persons) Unpack() []interface{} {
	res := make([]interface{}, 0)
	for _, i := range m {
		res = append(res, i)
	}
	return res
}

// Pack repackages a []interface into a Persons instance.
func (m Persons) Pack(replace []interface{}) Generic {
	res := make([]Person, 0)
	for _, value := range replace {
		res = append(res, value.(Person))
	}
	return Persons(res)
}

type PointerPersons []*Person

func (m PointerPersons) Unpack() []interface{} {
	res := make([]interface{}, 0)
	for _, i := range m {
		res = append(res, i)
	}
	return res
}

// Pack repackages a []interface into a Persons instance.
func (m PointerPersons) Pack(replace []interface{}) Generic {
	res := make([]*Person, 0)
	for _, value := range replace {
		res = append(res, value.(*Person))
	}
	return PointerPersons(res)
}

func TestIsEmpty(t *testing.T) {
	type args struct {
		s Slice
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test is empty string set",
			args: args{s: StringSlice{}},
			want: true,
		},
		{
			name: "test is not empty string set",
			args: args{s: StringSlice{"this", "is"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.s); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	testCache := []string{}

	type args struct {
		g        Generic
		modifier func(index int, val interface{})
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test ForEach Person",
			args: args{
				g: Persons{
					Person{Name: "billy"},
					Person{Name: "tommy"},
					Person{Name: "pat"},
				},
				modifier: func(index int, val interface{}) {
					item := val.(Person)
					testCache = append(testCache, item.Name)
				},
			},
			want: []string{"billy", "tommy", "pat"},
		},
		{
			name: "Test ForEach Person Empty",
			args: args{
				g: Persons{},
				modifier: func(index int, val interface{}) {
					item := val.(Person)
					testCache = append(testCache, item.Name)
				},
			},
			want: []string{},
		},
		{
			name: "Test ForEach PointerPerson",
			args: args{
				g: PointerPersons{
					&Person{Name: "billy"},
					&Person{Name: "tommy"},
					&Person{Name: "pat"},
				},
				modifier: func(index int, val interface{}) {
					item := val.(*Person)
					testCache = append(testCache, item.Name)
				},
			},
			want: []string{"billy", "tommy", "pat"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCache = []string{}

			ForEach(tt.args.g, tt.args.modifier)
			if !reflect.DeepEqual(testCache, tt.want) {
				t.Errorf("ForEach() = %v, want %v", testCache, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		g        Generic
		modifier func(index int, val interface{}) interface{}
	}
	tests := []struct {
		name string
		args args
		want Generic
	}{
		{
			name: "Test Map Person",
			args: args{
				g: Persons{
					Person{Name: "billy"},
					Person{Name: "tommy"},
					Person{Name: "pat"},
				},
				modifier: func(index int, val interface{}) interface{} {
					item := val.(Person)
					item.Name = item.Name + "_lol"
					return item
				},
			},
			want: Persons{
				Person{Name: "billy_lol"},
				Person{Name: "tommy_lol"},
				Person{Name: "pat_lol"},
			},
		},
		{
			name: "Test Map Person Empty",
			args: args{
				g: Persons{},
				modifier: func(index int, val interface{}) interface{} {
					item := val.(Person)
					item.Name = item.Name + "_lol"
					return item
				},
			},
			want: Persons{},
		},
		{
			name: "Test Map PointerPerson",
			args: args{
				g: PointerPersons{
					&Person{Name: "billy"},
					&Person{Name: "tommy"},
					&Person{Name: "pat"},
				},
				modifier: func(index int, val interface{}) interface{} {
					item := val.(*Person)
					item.Name = item.Name + "_lol"
					return item
				},
			},
			want: PointerPersons{
				&Person{Name: "billy_lol"},
				&Person{Name: "tommy_lol"},
				&Person{Name: "pat_lol"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.g, tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		g        Generic
		modifier func(index int, val interface{}) bool
	}
	tests := []struct {
		name string
		args args
		want Generic
	}{
		{
			name: "Test Filter Person",
			args: args{
				g: Persons{
					Person{Name: "billy"},
					Person{Name: "tommy"},
					Person{Name: "pat"},
				},
				modifier: func(index int, val interface{}) bool {
					item := val.(Person)
					if item.Name == "tommy" || item.Name == "billy" {
						return false
					}
					return true
				},
			},
			want: Persons{
				Person{Name: "pat"},
			},
		},
		{
			name: "Test Filter Person Empty",
			args: args{
				g: Persons{},
				modifier: func(index int, val interface{}) bool {
					item := val.(Person)
					if item.Name == "tommy" || item.Name == "billy" {
						return false
					}
					return true
				},
			},
			want: Persons{},
		},
		{
			name: "Test Filter PointerPerson",
			args: args{
				g: PointerPersons{
					&Person{Name: "billy"},
					&Person{Name: "tommy"},
					&Person{Name: "pat"},
				},
				modifier: func(index int, val interface{}) bool {
					item := val.(*Person)
					if item.Name == "tommy" || item.Name == "billy" {
						return false
					}
					return true
				},
			},
			want: PointerPersons{
				&Person{Name: "pat"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.g, tt.args.modifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
