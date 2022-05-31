package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
)

type Response struct {
	Success bool        `json:"success"`
	Error   error       `json:"error"`
	Msg     *string     `json:"msg"`
	Data    interface{} `json:"data"`
}

func CallSuccess0K(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Error:   nil,
		Msg:     &msg,
		Data:    data,
	})
	c.Next()

}

func CallServerError(c *gin.Context, msg string, err error){
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Error:   err,
		Msg:     &msg,
		Data:    map[string]interface{}{},
	})
	c.Abort()
}
