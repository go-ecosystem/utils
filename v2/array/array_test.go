package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContainString(t *testing.T) {
	type args struct {
		array []string
		item  string
	}

	var nilArray []string

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil array",
			args: args{
				array: nilArray,
				item:  "1",
			},
			want: false,
		},
		{
			name: "empty array",
			args: args{
				array: []string{},
				item:  "1",
			},
			want: false,
		},
		{
			name: "the string element is not in the array",
			args: args{
				array: []string{"1", "2", "3", "4", "5"},
				item:  "6",
			},
			want: false,
		},
		{
			name: "the string element is in the array",
			args: args{
				array: []string{"1", "2", "3", "4", "5"},
				item:  "1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.array, tt.args.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasDuplicateItem(t *testing.T) {
	type args struct {
		array []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				array: []string{
					"",
					"",
				},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				array: []string{
					"123",
					"",
					"ss",
					"123",
				},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				array: []string{
					"13",
					"",
					"ss",
					"123",
					"123asdasdqd",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := HasDuplicateItem(tt.args.array); got != tt.want {
				t.Errorf("HasDuplicateItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicateElements(t *testing.T) {
	array := []string{"1", "2", "3", "3", "1", "4", "5"}

	want := []string{"1", "2", "3", "4", "5"}

	if got := RemoveDuplicateElements(array); !reflect.DeepEqual(got, want) {
		t.Errorf("HasDuplicateItem() = %v, want %v", got, want)
	}
}

func TestMap(t *testing.T) {
	type args struct {
		data []int
		f    func(int) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "123",
			args: args{
				data: []int{1, 2, 3},
				f: func(i int) string {
					return fmt.Sprintf("%d", i)
				},
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.data, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
