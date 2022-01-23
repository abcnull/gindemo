package util

import (
	"gindemo/status"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// WrappedHandlerFunc 包装响应的入参，实际为路由函数
type WrappedHandlerFunc func(*gin.Context) (interface{}, *status.Status)

// WrappedResp 被包装的输出结构
type WrappedResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// WrapData 包装响应结果
func WrapData(handler WrappedHandlerFunc) func(*gin.Context) {
	return func(c *gin.Context) {
		data, stat := handler(c)

		// 如果 stat 不为空说明有 err 可以记录日志
		if stat != nil {
			log.Println("status 非空: ", stat)
		}

		// 被包装的响应结果 resp struct
		resp := new(WrappedResp)
		if stat != nil {
			resp.Code = stat.Code
			resp.Message = stat.Message
		} else {
			resp.Code = status.SUCCESS.Code
			resp.Message = status.SUCCESS.Message
		}
		resp.Data = data

		// 响应 json
		c.JSON(http.StatusOK, resp)
	}
}
