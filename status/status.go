package status

// Status 响应状态结构体
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	SUCCESS = &Status{0, "成功"}

	// 100xx 基本常用错误码
	REQ_PARAM_LACK_ERR    = &Status{10001, "请求缺少参数"}
	REQ_PARAM_INVALID_ERR = &Status{10002, "请求参数无效"}
	RPC_INVOKE_ERR        = &Status{10002, "远程调用出错"}
	INTERNAL_ERR          = &Status{10003, "内部错误"}
	TIMEOUT_ERR           = &Status{10004, "请求超时"}
	PAGE_NOT_FOUND        = &Status{10005, "页面找不到"}
	UNKNOW_ERR            = &Status{10006, "未知错误"}
	JSON_ERR              = &Status{10007, "JSON序列化错误"}
	DB_ERR                = &Status{10008, "数据库错误"}

	// 200xx 登录权限相关错误
	IDENTITY_EXPIRED = &Status{20001, "身份过期"}
	AUTH_FORBIDDEN   = &Status{20002, "没有权限"}

	// 310xx gindemo账号相关错误
	USER_NOT_EXIST = &Status{31001, "用户不存在"}
	// 311xx gindemo文章相关错误
	ARTICLE_NOT_EXIST = &Status{311001, "文章不存在"}

	// 320xx aaa平台错误

	// 330xx bbb平台错误
)
