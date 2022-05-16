package models

/* 유저의 모델 구조*/
type User struct {
	No        uint64 `json:"no"`
	Username  string `json:"user_name"`
	Password  string `json:"password"`
	UserEmail string `json:"email"`
}

/* Response  */
type UserResult struct {
	No        uint64 `json:"no"`
	Username  string `json:"user_name"`
	UserEmail string `json:"email"`
	CreatedAt string `json:"created_at"`
}

/* Request Body */
