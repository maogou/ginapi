package service

import (
	"context"
	"github.com/maogou/ginapi/app/dao"
	"github.com/maogou/ginapi/global"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

//实例化service实例
func New(ctx context.Context) Service  {
	service := Service{
		ctx: ctx,
	}

	service.dao = dao.New(global.DBEngine)

	return service
}
