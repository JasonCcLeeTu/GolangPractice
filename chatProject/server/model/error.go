package model

import(
	"errors"
)

var (

    USR_ERROR_NONEXIST = errors.New("此帳戶不存在")
	USR_ERROR_EXIST = errors.New("此帳戶已存在")
	USR_ERROR_PWD = errors.New("密碼錯誤")
    

)