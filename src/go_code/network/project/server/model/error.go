package model

import "errors"

var (
	ERROES_USER_NOTEXISTS = errors.New("用户不存在...")
	ERRORS_USER_EXISTS    = errors.New("用户已存在...")
	ERRORS_USER_PWD       = errors.New("密码不正确")
)
