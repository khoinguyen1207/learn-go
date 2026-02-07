package models

type User struct {
	ID    int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"column:name"`
	Email string `json:"email" gorm:"column:email"`
}
