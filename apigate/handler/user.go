package handler

import (
	"apigate/logic"
	"apigate/svc"
	"apigate/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

func RespErr(c *gin.Context, code int, msg string, err error) {
	if st, ok := status.FromError(err); ok {
		code = int(st.Code())
		msg = st.Message()
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func RespOK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": data,
	})
}

func CreateUser(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收请求参数
		var user types.UserInfo
		if err := c.ShouldBindJSON(&user); err != nil {
			zap.L().Error("c.ShouldBindJSON err: %v", zap.Error(err))
			RespErr(c, http.StatusBadRequest, "客户端请求错误", err)
			return
		}

		// 创建 logic 实例
		l := logic.NewUserSrvLogic(c.Request.Context(), svcCtx)
		// 调用 logic 方法
		resp, err := l.CreateUser(user)
		if err != nil {
			zap.L().Error("l.CreateUser err: %v", zap.Error(err))
			RespErr(c, http.StatusInternalServerError, "服务器内部错误", err)
			return
		}
		RespOK(c, resp)
	}
}
