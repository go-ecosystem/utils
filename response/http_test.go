package response

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewAPIModel(t *testing.T) {
	type args struct {
		code    Code
		message string
		data    interface{}
	}
	tests := []struct {
		name string
		args args
		want *APIModel
	}{
		{
			name: "",
			args: args{
				code:    CodeOK,
				message: "",
				data:    nil,
			},
			want: &APIModel{
				Code:    CodeOK,
				Data:    EmptyMapData(),
				Message: "",
			},
		},
		{
			name: "",
			args: args{
				code:    CodeOK,
				message: "asasa",
				data:    nil,
			},
			want: &APIModel{
				Code:    CodeOK,
				Data:    EmptyMapData(),
				Message: "asasa",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAPIModel(tt.args.code, tt.args.message, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPIModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	JSON(c, CodeOK, "test", nil)

	if w.Code != 200 {
		t.Errorf("JSON() error")
	}

	want := `{"code":0,"data":{},"message":"test"}`

	got := w.Body.String()
	if got != want {
		t.Errorf("JSON() = %v, want %v", got, want)
	}
}

func TestOK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	OK(c, "msg", "data")

	if w.Code != 200 {
		t.Errorf("OK() error")
	}

	want := `{"code":0,"data":"data","message":"msg"}`

	got := w.Body.String()
	if got != want {
		t.Errorf("OK() = %v, want %v", got, want)
	}
}

func TestError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Err(c, UnknownError(""))

	if w.Code != 200 {
		t.Errorf("Error() error")
	}

	want := `{"code":10000,"data":{},"message":"unknown error"}`

	got := w.Body.String()
	if got != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}

func TestEmptyData(t *testing.T) {
	m1 := EmptyMapData()
	m1["1"] = "1"

	m2 := EmptyMapData()
	if m2["1"] != nil {
		t.Errorf("Error() = %v, want %v", m2["1"], nil)
	}

	a1 := EmptyArrayData()
	a1 = append(a1, "1")

	a2 := EmptyArrayData()
	if len(a2) != 0 {
		t.Errorf("Error() = %v, want %v", len(a2), 0)
	}
}
