package core

import (
	"fmt"
	"strconv"
)

const CurrentRequesterKeyString = "CurrentRequesterKeyString"

type CurrentRequesterKeyStruct struct{}

type Requester interface {
	GetUserID() string
	GetUserIDInt() int

	GetRole() string
}

var _ Requester = (*RestRequester)(nil)

type RestRequester struct {
	ID   string
	Role string
}

func (u RestRequester) GetUserID() string {
	return u.ID
}

func (u RestRequester) GetUserIDInt() int {
	userIdInt, err := strconv.Atoi(u.ID)
	if err != nil {
		fmt.Println("Lỗi khi chuyển đổi userId sang int:", err)
		return -1
	}
	return userIdInt
}

func (u RestRequester) GetRole() string {
	return u.Role
}
