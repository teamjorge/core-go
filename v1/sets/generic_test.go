package sets

import (
	"errors"
	"reflect"
	"sort"
	"testing"
)

/*
******************************************
				Mock Set
******************************************
*/

type User struct {
	Name string
	Age  int
}

type UserSet map[User]bool

func (i UserSet) Add(in ...interface{}) {
	for _, x := range in {
		i[x.(User)] = false
	}
}

func (i UserSet) Delete(elems ...interface{}) {
	for _, x := range elems {
		delete(i, x.(User))
	}
}

func (i UserSet) ToSlice() interface{} {
	res := make([]User, 0)
	for key := range i {
		res = append(res, key)
	}
	return res
}

/*
******************************************
******************************************
******************************************
 */

func TestIsEmpty(t *testing.T) {
	type args struct {
		s Set
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test is empty string set",
			args: args{s: NewString()},
			want: true,
		},
		{
			name: "test is not empty string set",
			args: args{s: NewString([]string{"this", "is"})},
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

func Test_resolveError(t *testing.T) {
	type args struct {
		r interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "test resolveError is string",
			args: args{
				r: "cashmoney",
			},
			wantErr: true,
			want:    "cashmoney",
		},
		{
			name: "test resolveError is error",
			args: args{
				r: errors.New("cashmoney"),
			},
			wantErr: true,
			want:    "cashmoney",
		},
		{
			name: "test resolveError is something else",
			args: args{
				r: map[string]string{"error": "cashmoney"},
			},
			wantErr: true,
			want:    "map[error:cashmoney]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := resolveError(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("resolveError() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if !reflect.DeepEqual(err.Error(), tt.want) {
					t.Errorf("resolveError() error = %s, want %s", err.Error(), tt.want)
				}
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		g     Generic
		items []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    Generic
	}{
		{
			name: "test add no error",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
				},
				items: []interface{}{
					User{Name: "billy", Age: 9},
					User{Name: "bobby", Age: 12},
					User{Name: "tommy", Age: 15},
					User{Name: "jane", Age: 17},
					User{Name: "jane", Age: 19},
				},
			},
			wantErr: false,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
				User{Name: "tommy", Age: 15}:  false,
				User{Name: "jane", Age: 17}:   false,
				User{Name: "jane", Age: 19}:   false,
			},
		},
		{
			name: "test add with error",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
				},
				items: []interface{}{
					User{Name: "billy", Age: 9},
					User{Name: "bobby", Age: 12},
					User{Name: "tommy", Age: 15},
					"I'm totally a User",
					User{Name: "jane", Age: 17},
					User{Name: "jane", Age: 19},
				},
			},
			wantErr: true,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
			},
		},
		{
			name: "test add empty",
			args: args{
				g:     UserSet{},
				items: []interface{}{},
			},
			wantErr: false,
			want:    UserSet{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Add(tt.args.g, tt.args.items...); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if !reflect.DeepEqual(tt.args.g, tt.want) {
					t.Errorf("Add() g = %v, want %v", tt.args.g, tt.want)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		g     Generic
		items []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    Generic
	}{
		{
			name: "test delete no error",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
					User{Name: "tommy", Age: 15}:  false,
					User{Name: "jane", Age: 17}:   false,
					User{Name: "jane", Age: 19}:   false,
				},
				items: []interface{}{
					User{Name: "george", Age: 12},
					User{Name: "jane", Age: 17},
				},
			},
			wantErr: false,
			want: UserSet{
				User{Name: "billy", Age: 9}:  false,
				User{Name: "bobby", Age: 12}: false,
				User{Name: "tommy", Age: 15}: false,
				User{Name: "jane", Age: 19}:  false,
			},
		},
		{
			name: "test delete key does not exist",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
					User{Name: "tommy", Age: 15}:  false,
					User{Name: "jane", Age: 17}:   false,
					User{Name: "jane", Age: 19}:   false,
				},
				items: []interface{}{
					User{Name: "franky", Age: 5},
				},
			},
			wantErr: false,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
				User{Name: "tommy", Age: 15}:  false,
				User{Name: "jane", Age: 17}:   false,
				User{Name: "jane", Age: 19}:   false,
			},
		},
		{
			name: "test delete with error",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
					User{Name: "tommy", Age: 15}:  false,
					User{Name: "jane", Age: 17}:   false,
					User{Name: "jane", Age: 19}:   false,
				},
				items: []interface{}{
					User{Name: "billy", Age: 9},
					User{Name: "bobby", Age: 12},
					User{Name: "tommy", Age: 15},
					"I'm totally a User",
					User{Name: "jane", Age: 17},
					User{Name: "jane", Age: 19},
				},
			},
			wantErr: true,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
				User{Name: "tommy", Age: 15}:  false,
				User{Name: "jane", Age: 17}:   false,
				User{Name: "jane", Age: 19}:   false,
			},
		},
		{
			name: "test delete empty",
			args: args{
				g:     UserSet{},
				items: []interface{}{},
			},
			wantErr: false,
			want:    UserSet{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Delete(tt.args.g, tt.args.items...); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if !reflect.DeepEqual(tt.args.g, tt.want) {
					t.Errorf("Delete() g = %v, want %v", tt.args.g, tt.want)
				}
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	type args struct {
		g Generic
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "test to slice",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "tommy", Age: 15}:  false,
				},
			},
			want: []User{
				{Name: "george", Age: 12},
				{Name: "tommy", Age: 15},
			},
			wantErr: false,
		},
		{
			name: "test to slice same name",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "george", Age: 15}: false,
					User{Name: "tommy", Age: 15}:  false,
				},
			},
			want: []User{
				{Name: "george", Age: 12},
				{Name: "george", Age: 15},
				{Name: "tommy", Age: 15},
			},
			wantErr: false,
		},
		{
			name:    "test to slice empty",
			args:    args{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := ToSlice(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if res == nil {
				return
			}
			got := res.([]User)
			sort.SliceStable(got, func(i, j int) bool {
				if got[i].Name == got[j].Name {
					return got[i].Age < got[j].Age
				}
				return got[i].Name < got[j].Name
			})
			if got != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rollbackAdd(t *testing.T) {
	type args struct {
		g     Generic
		items []interface{}
	}
	tests := []struct {
		name string
		args args
		want Generic
	}{
		{
			name: "test rollback add",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
					User{Name: "tommy", Age: 15}:  false,
					User{Name: "jane", Age: 17}:   false,
					User{Name: "jane", Age: 19}:   false,
				},
				items: []interface{}{
					User{Name: "tommy", Age: 15},
					User{Name: "jane", Age: 17},
					User{Name: "jane", Age: 19},
				},
			},
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rollbackAdd(tt.args.g, tt.args.items...)
			if tt.args.g != nil && !reflect.DeepEqual(tt.args.g, tt.want) {
				t.Errorf("rollbackAdd() = %v, want %v", tt.args.g, tt.want)
			}
		})
	}
}

func TestAddUnsafe(t *testing.T) {
	type args struct {
		g     Generic
		items []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    Generic
	}{
		{
			name: "test unsafe add no error",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
				},
				items: []interface{}{
					User{Name: "tommy", Age: 15},
					User{Name: "jane", Age: 17},
					User{Name: "jane", Age: 19},
				},
			},
			wantErr: false,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
				User{Name: "tommy", Age: 15}:  false,
				User{Name: "jane", Age: 17}:   false,
				User{Name: "jane", Age: 19}:   false,
			},
		},
		{
			name: "test unsafe add with error",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
				},
				items: []interface{}{
					User{Name: "tommy", Age: 15},
					"im totes a user",
					User{Name: "jane", Age: 17},
					User{Name: "jane", Age: 19},
				},
			},
			wantErr: true,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
				User{Name: "tommy", Age: 15}:  false,
			},
		},
		{
			name: "test unsafe add wrong",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
				},
				items: []interface{}{
					"im totes a user",
				},
			},
			wantErr: true,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
			},
		},
		{
			name: "test unsafe add nothing",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
				},
				items: []interface{}{},
			},
			wantErr: false,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddUnsafe(tt.args.g, tt.args.items...); (err != nil) != tt.wantErr {
				t.Errorf("AddUnsafe() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if tt.args.g != nil && !reflect.DeepEqual(tt.args.g, tt.want) {
					t.Errorf("AddUnsafe() = %v, want %v", tt.args.g, tt.want)
				}
			}
		})
	}
}

func Test_rollbackDelete(t *testing.T) {
	type args struct {
		g     Generic
		items []interface{}
	}
	tests := []struct {
		name string
		args args
		want Generic
	}{
		{
			name: "test rollback delete",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
					User{Name: "tommy", Age: 15}:  false,
					User{Name: "jane", Age: 17}:   false,
					User{Name: "jane", Age: 19}:   false,
				},
				items: []interface{}{
					User{Name: "tommy", Age: 15},
					User{Name: "jane", Age: 17},
					User{Name: "jane", Age: 19},
				},
			},
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
				User{Name: "tommy", Age: 15}:  false,
				User{Name: "jane", Age: 17}:   false,
				User{Name: "jane", Age: 19}:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rollbackDelete(tt.args.g, tt.args.items...)
			if tt.args.g != nil && !reflect.DeepEqual(tt.args.g, tt.want) {
				t.Errorf("rollbackDelete() = %v, want %v", tt.args.g, tt.want)
			}
		})
	}
}

func TestDeleteUnsafe(t *testing.T) {
	type args struct {
		g     Generic
		items []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    Generic
	}{
		{
			name: "test unsafe delete no error",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
					User{Name: "tommy", Age: 15}:  false,
					User{Name: "jane", Age: 17}:   false,
					User{Name: "jane", Age: 19}:   false,
				},
				items: []interface{}{
					User{Name: "tommy", Age: 15},
					User{Name: "jane", Age: 17},
					User{Name: "jane", Age: 19},
				},
			},
			wantErr: false,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
			},
		},
		{
			name: "test unsafe delete key does not exist",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
					User{Name: "tommy", Age: 15}:  false,
					User{Name: "jane", Age: 17}:   false,
					User{Name: "jane", Age: 19}:   false,
				},
				items: []interface{}{
					User{Name: "franky", Age: 20},
					User{Name: "jane", Age: 17},
					User{Name: "jane", Age: 19},
				},
			},
			wantErr: false,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
				User{Name: "tommy", Age: 15}:  false,
			},
		},
		{
			name: "test unsafe delete with error",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
					User{Name: "tommy", Age: 15}:  false,
					User{Name: "jane", Age: 17}:   false,
					User{Name: "jane", Age: 19}:   false,
				},
				items: []interface{}{
					User{Name: "tommy", Age: 15},
					User{Name: "jane", Age: 17},
					"im totes a user",
					User{Name: "jane", Age: 19},
				},
			},
			wantErr: true,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
				User{Name: "jane", Age: 19}:   false,
			},
		},
		{
			name: "test unsafe delete wrong",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
					User{Name: "tommy", Age: 15}:  false,
					User{Name: "jane", Age: 17}:   false,
					User{Name: "jane", Age: 19}:   false,
				},
				items: []interface{}{
					"im totes a user",
				},
			},
			wantErr: true,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
				User{Name: "tommy", Age: 15}:  false,
				User{Name: "jane", Age: 17}:   false,
				User{Name: "jane", Age: 19}:   false,
			},
		},
		{
			name: "test unsafe delete nothing",
			args: args{
				g: UserSet{
					User{Name: "george", Age: 12}: false,
					User{Name: "billy", Age: 9}:   false,
					User{Name: "bobby", Age: 12}:  false,
					User{Name: "tommy", Age: 15}:  false,
					User{Name: "jane", Age: 17}:   false,
					User{Name: "jane", Age: 19}:   false,
				},
				items: []interface{}{},
			},
			wantErr: false,
			want: UserSet{
				User{Name: "george", Age: 12}: false,
				User{Name: "billy", Age: 9}:   false,
				User{Name: "bobby", Age: 12}:  false,
				User{Name: "tommy", Age: 15}:  false,
				User{Name: "jane", Age: 17}:   false,
				User{Name: "jane", Age: 19}:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUnsafe(tt.args.g, tt.args.items...); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUnsafe() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if tt.args.g != nil && !reflect.DeepEqual(tt.args.g, tt.want) {
					t.Errorf("DeleteUnsafe() = %v, want %v", tt.args.g, tt.want)
				}
			}
		})
	}
}
