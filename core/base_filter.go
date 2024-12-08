package core

type BaseFilter struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
