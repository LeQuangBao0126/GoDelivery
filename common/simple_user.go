package common




type SimpleUser struct {
	SQLModel `json:",inline"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:roles;"`
}

func (u *SimpleUser) GetUserId() int {
	return u.Id
}

func (u *SimpleUser) GetRole() string {
	return u.Role
}

func (SimpleUser) TableName() string {
	return "users"
}
