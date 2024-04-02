package abc

import "errors"

var (
	UserConfligct = errors.New("người dùng không tồn tại")
	UserNotFoind  = errors.New("ngươid dùng không tồn tại")
	SignUpFail    = errors.New("Thất bại")
)
