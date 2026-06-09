// 定义了 http 请求和响应结构体，负责与 http 客户端交互
package types

type UserInfo struct {
	UserID   uint32 `json:"user_id"`
	NickName string `json:"nick_name"`
	Password string `json:"password"`
	Mobile   string `josn:"mobile"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Gender   uint8  `json:"gender"`
	Role     uint8  `json:"role"`
}
