package model

// User used to
type User struct {
	Eamil    string `json:"email"`
	Password string `json:"password"`
}

// UserInfo used to
type UserInfo struct {
	User
	Name string `json:"name"`
	Role string `json:"role"`
}
