package http_response

import (
	"Managing-home-energy/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code       string `json:"code"`
	DebugStack string `json:"debug_stack"`
	Message    string `json:"message"`
	Version    string `json:"version"`
}

type ResultResponse struct { // embed struct into struct
	Response             // struct embed
	Result   interface{} `json:"result"` // is one field in struct, accept all kind of value
}

func ReturnErrMessage(c *gin.Context, responseCode string, err error) {
	var empty struct{}
	r := ResultResponse{
		Response: Response{
			Code:       responseCode,
			DebugStack: err.Error(),
			Message:    GetMessage(responseCode),
			Version:    "v1",
		},
		Result: empty,
	}
	c.Set(constants.RequestErrorMessage, GetMessage(responseCode))
	c.AbortWithStatusJSON(GetHttpResponseStatus(responseCode), r)
}

func ReturnSuccessMessage(c *gin.Context, data interface{}) {
	if data == nil {
		var empty struct{}
		data = empty
	}
	result := ResultResponse{
		Response: Response{
			Code:       ResponseOK,
			Message:    GetMessage(ResponseOK),
			Version:    "v1",
			DebugStack: "",
		},
		Result: data,
	}
	c.Set(constants.RequestErrorMessage, GetMessage(ResponseOK))
	c.JSON(http.StatusOK, result)
}
