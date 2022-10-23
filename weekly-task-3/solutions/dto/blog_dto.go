package dto

type BlogDTO struct {
	UserId     string `json:"user_id" form:"user_id"`
	CategoryId string `json:"category_id" form:"category_id"`
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
}
