package handler

import (
	"apigate/logic"
	"apigate/svc"
	"apigate/types"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RespErr(c *gin.Context, code int, msg string, err error) {
	if st, ok := status.FromError(err); ok && st.Code() != codes.OK {
		code = int(st.Code())
		msg = st.Message()
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": code,
		"msg":  msg,
	})
}

// Captcha 获取验证码请求
func Captcha(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 调用logic
		l := logic.NewUserSrvLogic(c.Request.Context(), svcCtx)
		resp, err := l.Captcha()
		if err != nil {
			zap.L().Error("l.Captcha err: %v", zap.Error(err))
			RespErr(c, http.StatusInternalServerError, "服务端内部错误", err)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// Register 注册

// Login 用户登录
func Login(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求参数
		var req types.LoginReq
		if err := c.ShouldBindJSON(&req); err != nil {
			zap.L().Error("c.ShouldBindJSON err: %v", zap.Error(err))
			RespErr(c, http.StatusBadRequest, "客户端请求错误", err)
			return
		}
		// 创建logic实例
		l := logic.NewUserSrvLogic(c.Request.Context(), svcCtx)
		resp, err := l.Login(req)
		if err != nil {
			zap.L().Error("l.Login err: %v", zap.Error(err))
			RespErr(c, http.StatusInternalServerError, "服务器内部错误", err)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// CreateUser 创建新用户，只有管理员才有权限做
func CreateUser(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建 logic 实例
		l := logic.NewUserSrvLogic(c.Request.Context(), svcCtx)

		// 验证当前用户权限
		userID, exists := c.Get("userID")
		if !exists {
			RespErr(c, http.StatusUnauthorized, "未认证", nil)
			return
		}
		id, ok := userID.(uint32)
		if !ok {
			RespErr(c, http.StatusInternalServerError, "服务器内部错误", errors.New("当前登录用户id格式不正确"))
			return
		}
		admin, err := l.GetUserByID(types.GetUserByID{UserID: id})
		if err != nil {
			RespErr(c, http.StatusInternalServerError, "服务器内部错误", err)
			return
		}

		adminData, ok := admin.Data.(*types.UserInfo)
		if !ok {
			RespErr(c, http.StatusInternalServerError, "服务器内部错误", errors.New("断言失败"))
			return
		}
		if adminData.Role != 2 {
			RespErr(c, http.StatusForbidden, "禁止访问", nil)
			return
		}

		// 接收请求参数
		var user types.UserInfo
		if err := c.ShouldBindJSON(&user); err != nil {
			zap.L().Error("c.ShouldBindJSON err: %v", zap.Error(err))
			RespErr(c, http.StatusBadRequest, "客户端请求错误", err)
			return
		}
		// 调用 logic 方法
		resp, err := l.CreateUser(user)
		if err != nil {
			zap.L().Error("l.CreateUser err: %v", zap.Error(err))
			RespErr(c, http.StatusInternalServerError, "服务器内部错误", err)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// GetUserByID 根据id获取用户信息
func GetUserByID(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收请求参数
		var req types.GetUserByID
		if err := c.ShouldBindJSON(&req); err != nil {
			RespErr(c, http.StatusBadRequest, "客户端请求错误", err)
			return
		}
		// 调用logic层
		l := logic.NewUserSrvLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetUserByID(req)
		if err != nil {
			RespErr(c, http.StatusInternalServerError, "服务器内部错误", err)
		}
		c.JSON(http.StatusOK, resp)
	}
}
