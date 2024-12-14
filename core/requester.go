package core

const CurrentRequesterKeyString = "CurrentRequesterKeyString"

type CurrentRequesterKeyStruct struct{}

type Requester interface {
	GetUserID() string
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

func (u RestRequester) GetRole() string {
	return u.Role
}
