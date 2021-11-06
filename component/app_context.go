package component

import (
	"ProjectDelivery/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() (*gorm.DB)
	SecretKey() string
	GetPubsub() pubsub.Pubsub
}
type appCtx struct {
	db *gorm.DB
	secretKey string
	pb pubsub.Pubsub
}
func NewAppContext(db *gorm.DB ,secretKey string ,pb pubsub.Pubsub) *appCtx{
	 return &appCtx{db : db , secretKey: secretKey ,pb :pb}
}
func (a *appCtx) GetMainDBConnection() *gorm.DB{
	return  a.db
}
func (a *appCtx) SecretKey()string {
	return a.secretKey
}
func (ctx *appCtx) GetPubsub()  pubsub.Pubsub {
	 return  ctx.pb
}