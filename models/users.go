package models

type User struct {
	// gorm.Model
	Id       int      `json:"id"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Message  string   `json: "message"`
	Results  []Result `gorm:"foreignKey:UserID"`
}

type UserLogin struct {
	Id       int    `json:"id" db="id"`
	Email    string `json:"email" db="email"`
	Password string `json:"password" db="password"`
}

type UserLoginReq struct {
	Email    string `json:"email" db="email"`
	Password string `json:"password" db="password"`
}

// type UserLoginResp struct{

// }
