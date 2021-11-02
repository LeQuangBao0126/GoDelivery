package component

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
}
type appCtx struct {
	db *gorm.DB
}
func NewAppContext(db *gorm.DB) *appCtx{
	 return &appCtx{db : db}
}
func (a *appCtx) GetMainDBConnection() *gorm.DB{
	return  a.db
}