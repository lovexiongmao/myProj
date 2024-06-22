package models

type AccountInfo struct {
	ID        uint   `json:"-" gorm:"primaryKey"`
	Username  string `json:"username" `
	Password  string `json:"password"`
	Role      string `json:"role"`
	BaseModel `json:"-"`
}
