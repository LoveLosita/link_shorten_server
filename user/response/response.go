package response

import "link_shorten_server/user/kitex_gen/user"

func InternalErr(err error) user.Status {
	return user.Status{Code: "500", Message: err.Error()}
}

var (
	Ok = user.Status{
		Code:    "10000",
		Message: "ok",
	}
	WrongUsrName = user.Status{
		Code:    "40001",
		Message: "wrong username",
	}
	UsernameExists = user.Status{
		Code:    "40002",
		Message: "the username already exists",
	}
	WrongGender = user.Status{
		Code:    "40003",
		Message: "wrong gender",
	}
)
