package dao

import (
	"github.com/maogou/ginapi/app/model"
	"github.com/maogou/ginapi/pkg/app"
)

//新增tag
func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}

	return tag.Create(d.engine)
}

//标签总数
func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}

	return tag.Count(d.engine)
}

//标签列表
func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}

	pageOffset := app.GetPageOffset(page, pageSize)

	return tag.List(d.engine, pageOffset, pageSize)
}

//更新tag
func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}

	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}

	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

//删除标签
func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}

	return tag.Delete(d.engine)
}
