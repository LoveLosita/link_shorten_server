package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"link_shorten_server/user/kitex_gen/user"
	"link_shorten_server/user/response"
)

var accessKey = AccessKey

func CheckJwtToken(token string) (int, user.Status) {
	//fmt.Println(string(tokenString))//测试用
	if string(token) == "" { //没有token
		return -1, response.MissingToken
	}
	// 解析并验证 Token
	token2, err := jwt.Parse(string(token), func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法是我们支持的 HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return accessKey, nil
	})
	if err != nil || !token2.Valid { //token无效
		return -1, response.InvalidToken
	}

	// 将解析出的用户信息存入上下文，供后续使用
	if claims, ok := token2.Claims.(jwt.MapClaims); ok {
		// 获取 token_type 判断类型
		tokenType, ok := claims["token_type"].(string)
		if !ok {
			return -1, response.InvalidClaims
		}
		// 根据 token_type 做不同的处理
		if tokenType == "access_token" {
			// 如果是访问令牌，可以设置用户ID并继续
			/*c.Set("user_id", claims["user_id"])
			return*/
			floatID := claims["user_id"].(float64)
			return int(floatID), user.Status{}
		} else {
			return -1, response.WrongTokenType
		}
	} else {
		return -1, response.InvalidToken
	}
}
