package dao

import "github.com/maogou/ginapi/app/model"

func (d *Dao) GetAuth(appKey,appSecret string) (model.Auth,error)  {
	auth := model.Auth{
		AppKey: appKey,
		AppSecret: appSecret,
	}

	return auth.Get(d.engine)
}


