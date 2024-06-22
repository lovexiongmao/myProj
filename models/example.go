package models

type User struct {
	ID        uint   `json:"-" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"column:name"`
	Age       int    `json:"age" gorm:"column:age"`
	Gender    string `json:"gender" gorm:"column:gender"`
	Height    int    `json:"height" gorm:"column:height"`
	BaseModel `json:"-"`
}
