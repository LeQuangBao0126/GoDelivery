package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Status    int        `json:"status" gorm:"status"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
}

//func (m *SQLModel ) GenUID(dbType int){
//	uid := NewUID(uint32(m.Id),dbType,1)
//	m.FakeId = &uid
//}