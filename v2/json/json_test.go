package json

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "data is nil, return empty string",
			args: args{
				data: nil,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "data is a string",
			args: args{
				data: "这是一个测试",
			},
			want:    `"这是一个测试"`,
			wantErr: false,
		},
		{
			name: "data is number",
			args: args{
				data: 1000,
			},
			want:    "1000",
			wantErr: false,
		},
		{
			name: "data is an array",
			args: args{
				data: []int{1, 2, 3},
			},
			want:    "[1,2,3]",
			wantErr: false,
		},
		{
			name: "data is a map",
			args: args{
				data: map[string]int{"k1": 1, "k2": 2},
			},
			want:    `{"k1":1,"k2":2}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToString(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMap(t *testing.T) {
	type model struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "data is nil",
			args: args{
				data: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "data is a map",
			args: args{
				data: `{"k1":1}`,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "data is a string",
			args: args{
				data: model{
					ID:   1,
					Name: "name1",
				},
			},
			want: map[string]interface{}{
				"id":   1,
				"name": "name1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToMap(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("ToMap() = %v, want %v", got, tt.want)
		})
	}
}

func TestConvert(t *testing.T) {
	type args struct {
		s interface{}
		d interface{}
	}
	slice := make([]int, 0)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "array slice",
			args: args{
				s: []int{1, 2, 3},
				d: &slice,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Convert(tt.args.s, tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("Convert() = %v, to %v \n", tt.args.s, slice)
		})
	}
}

func TestFromString(t *testing.T) {
	type args struct {
		str  string
		data interface{}
	}
	slice := make([]int, 0)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no json string",
			args: args{
				str: "这是一个测试字符串",
			},
			wantErr: true,
		},
		{
			name: "aray string",
			args: args{
				str:  "[1,2,3]",
				data: &slice,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FromString(tt.args.str, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("FromString() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Printf("FromString() str = %v, to = %v \n", tt.args.str, slice)
		})
	}
}
