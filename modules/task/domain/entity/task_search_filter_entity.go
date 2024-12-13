package entity

type TaskSearchFilterEntity struct {
	UserID   *int    `form:"userId" binding:"omitempty,min=1"`
	Status   *string `form:"status" binding:"omitempty,oneof=Todo InProgress Completed Expired"`
	Priority *string `form:"priority" binding:"omitempty,oneof=Low Medium High"`
	Search   *string `form:"search" binding:"omitempty"`
	SortBy   *string `form:"sortBy" binding:"omitempty,oneof=taskName priority status dueDate"`
	Order    *string `form:"order" binding:"omitempty,oneof=asc desc"`
	Limit    *int    `form:"limit" binding:"omitempty,min=1"`
	Page     *int    `form:"page" binding:"omitempty,min=1"`
}
