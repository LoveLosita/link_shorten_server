package utils

import (
	"strconv"
)

func IDToAbc(id int) (string, error) { //用于生成短链
	if id < 0 {
		return "", strconv.ErrSyntax
	}
	charSet := "abcdefghij"
	var result string
	var last int
	for id > 0 {
		last = id % 10
		id = id / 10
		result = string(charSet[last]) + result
	}
	return result, nil
}
