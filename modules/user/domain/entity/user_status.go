package entity

type UserStatus int

const (
	ACTIVE UserStatus = iota + 1
	BLOCKED
	UNVERIFIED
)

func (status UserStatus) Value() string {
	switch status {
	case ACTIVE:
		return "ACTIVE"
	case BLOCKED:
		return "BLOCKED"
	case UNVERIFIED:
		return "UNVERIFIED"
	}

	return ""
}
