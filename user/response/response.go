package response

import "link_shorten_server/user/kitex_gen/user"

func InternalErr(err error) user.Status {
	return user.Status{Code: "500", Message: err.Error()}
}

func ResponseTo(err error) user.Status {
	if err == nil {
		return user.Status{}
	}
	return user.Status{Code: "500", Message: err.Error()}
}

var (
	Ok = user.Status{ //正常,和客户端统一
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
	EmailExists = user.Status{
		Code:    "40004",
		Message: "the email already exists",
	}
	PhoneNumberExists = user.Status{
		Code:    "40005",
		Message: "the phone number already exists",
	}
	WrongPassword = user.Status{
		Code:    "40006",
		Message: "wrong password",
	}
	InvalidTokenSingingMethod = user.Status{
		Code:    "40007",
		Message: "invalid token signing method",
	}
	InvalidClaims = user.Status{
		Code:    "40008",
		Message: "invalid claims",
	}
	WrongTokenType = user.Status{
		Code:    "40009",
		Message: "wrong token type",
	}
	MissingToken = user.Status{
		Code:    "40010",
		Message: "missing token",
	}
	InvalidToken = user.Status{
		Code:    "40011",
		Message: "invalid token",
	}
)
