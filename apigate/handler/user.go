package handler

import (
	"apigate/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func rpcErrResp(err error, c *gin.Context) {
	if sta, ok := status.FromError(err); ok {
		switch sta.Code() {
		case codes.NotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "请求资源不存在",
			})
		case codes.Internal:
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "服务器内部错误",
			})
		case codes.InvalidArgument:
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "请求参数错误",
			})
		case codes.Unauthenticated:
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未登录",
			})
		case codes.PermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{
				"code": 403,
				"msg":  "禁止访问",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "未知错误",
			})
		}
	}
}
func CreateUser(c *gin.Context) {
	// 接收请求参数
	var user types.UserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "客户端请求错误",
		})
	}

}
