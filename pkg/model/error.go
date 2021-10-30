package model

import "errors"

type (
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
)

var (
	ERR_USER_NOT_FOUND = errors.New("ERR_USER_NOT_FOUND")
)
