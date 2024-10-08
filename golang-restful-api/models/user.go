package models

//Structur Table User
type User struct{
	
	Id int64 `gorm:"primary_key"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
	Data []Transparasi `gorm:"foreignKey:User_id"`

}