package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Username string `json:"username" gorm:"type:varchar(100);unique"`
	Level    string `json:"level" gorm:"type:int(10)"`
	Password []byte `json:"-" gorm:"type:varchar(100)"`
}
