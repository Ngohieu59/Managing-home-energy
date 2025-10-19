package dto

import "errors"

var ErrUserAlreadyExists = errors.New("username already exist")

var ErrUserIDNotFound = errors.New("not found user with that ID")

var ErrUserNameNotFound = errors.New("not found user with that name")

var ErrNotFillAllFields = errors.New("please fill all fields")

var ErrDataFormatWrong = errors.New("wrong data format")

var ErrCreateUserMysql = errors.New("fail to create user to Mysql")

var ErrUserTokenInvalidOrExpired = errors.New("token is invalid or already expired")

var ErrPermissionDenied = errors.New("this user does not have permission to do this")

var ErrDeleteUserMysql = errors.New("fail to delete user to Mysql")

var ErrStrconv = errors.New("fail convert string to int")

var ErrPasswordIncorrect = errors.New("password incorrect")

var ErrStartAfterEnd = errors.New("start date after end date")
