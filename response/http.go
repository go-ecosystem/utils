package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PageModel page response model
type PageModel struct {
	List       interface{} `json:"list"`
	Number     int         `json:"pageNumber"`
	Size       int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
}

// APIModel api response model
type APIModel struct {
	Code    Code        `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// NewAPIModel new api response
func NewAPIModel(code Code, message string, data interface{}) *APIModel {
	if data == nil {
		data = EmptyMapData
	}
	res := &APIModel{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return res
}

// EmptyMapData EmptyMapData
var EmptyMapData = make(map[string]interface{})

// EmptyArrayData EmptyArrayData
var EmptyArrayData = make([]interface{}, 0)

// JSON response APIModel
func JSON(c *gin.Context, code Code, message string, data interface{}) {
	res := NewAPIModel(code, message, data)
	c.JSON(http.StatusOK, res)
}

// OK response OK
func OK(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, NewAPIModel(CodeOK, message, data))
}

// Err response error
func Err(c *gin.Context, err Error) {
	res := NewAPIModel(err.Code, err.Error(), EmptyMapData)
	c.JSON(http.StatusOK, res)
}

// Abort response abort
func Abort(c *gin.Context, err Error) {
	Err(c, err)
	c.Abort()
}
