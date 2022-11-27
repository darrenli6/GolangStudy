package dao

import (
	"github.com/darrenli6/blog-server/internal/model"
	"github.com/darrenli6/blog-server/pkg/app"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}

	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {

	tag := model.Tag{Name: name, State: state}
	pageOffSet := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffSet, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreateBy: createdBy},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modified string) error {

	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}

	values := map[string]interface{}{
		"state":       state,
		"modified_by": modified,
	}

	if name != "" {
		values[name] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}
