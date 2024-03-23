package requests

import "time"

type GetUserInfoResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	City     string `json:"city"`
}

type ShowHistory struct {
	CreateTime time.Time `json:"create_time"`
	FileName   string    `json:"file_name"`
	TotalScore float64   `json:"total_score"`
}
