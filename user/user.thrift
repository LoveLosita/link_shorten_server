namespace go user

struct Status {
  1: required string code
  2: required string message
}

struct UserRegisterRequest { #用户注册请求
  1: required string username
  2: required string password
  3: required string gender
  4: required string phone_number
  5: required string email
}

struct UserRegisterResponse { #用户注册响应
  1: required Status status
}

struct UserLoginRequest { #用户登录请求
  1: required string username
  2: required string password
}

struct UserLoginResponseData { #用户登录响应数据
  1: optional string access_token
  2: optional string refresh_token
}

struct UserLoginResponse { #用户登录响应
  1: required Status status
  2: optional UserLoginResponseData data #用户登录响应数据
}

struct TokenRefreshRequest { #刷新token请求
  1: required string refresh_token
}

struct TokenRefreshResponseData { #刷新token响应数据
  1: required string access_token
  2: required string refresh_token
}

struct TokenRefreshResponse { #刷新token响应
  1: required Status status
  2: optional TokenRefreshResponseData data #刷新token响应数据
}

service UserService {
  UserRegisterResponse user_register(1: UserRegisterRequest req) #用户注册
  UserLoginResponse user_login(1: UserLoginRequest req) #用户登录
  TokenRefreshResponse token_refresh(1: TokenRefreshRequest req) #刷新token
}