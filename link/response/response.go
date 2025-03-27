package response

import (
	"link_shorten_server/link/kitex_gen/link"
)

func InternalErr(err error) link.Status {
	return link.Status{Code: "500", Message: err.Error()}
}

var (
	Ok = link.Status{ //正常,和客户端统一
		Code:    "10000",
		Message: "ok",
	}
	EmptyRequest = link.Status{ //请求为空
		Code:    "41001",
		Message: "request is empty",
	}
	LinkNotExists = link.Status{ //链接不存在
		Code:    "41002",
		Message: "link not exists",
	}
)
