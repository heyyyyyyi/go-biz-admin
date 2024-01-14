package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"` // 返回的请求体里不包括 will not return this field if add -
	RoleId    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}

// 成员函数
func (user *User) TranslatePassword(password string) {
	temp, _ = bcrypt.GenerateFromPassword([]byte(password), 14) // 14?
	User.Password = string(temp)
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password))
}
