package persistence

import (
	"errors"
	"strings"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

func NewAttribute(db *xorm.Engine) repository.Attribute {
	return &attribute{db: db}
}

// 商品属性
type attribute struct {
	db *xorm.Engine
}

// 创建属性
func (t *attribute) Create(info po.Attribute) error {
	has, err := t.db.UseBool("multi,required,is_sale").Insert(&info)
	if err != nil || has == 0 {
		return errors.New("写入失败")
	}
	return nil
}

// 删除属性
func (t *attribute) Delete(id []int64) error {
	info := new(po.Attribute)
	has, err := t.db.In("id", id).Delete(info)
	if err != nil || has == 0 {
		return errors.New("删除失败")
	}

	return nil
}

// 更新属性
func (t *attribute) Update(info po.Attribute) error {
	has, err := t.db.Where("id=?", info.ID).UseBool("multi,required,is_sale").Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

// 属性列表
func (t *attribute) Find(req query.Request) (list []po.Attribute, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "label":
			query += "label LIKE ?"
			args = append(args, v.Value.(string)+"%")
		case "type":
			query += " AND type=?"
			args = append(args, v.Value)
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	info := new(po.Attribute)

	// 统计条数
	total, err = t.db.Where(query, args...).Count(info)
	if err != nil {
		return
	}
	if total > 0 {
		// 分页数据
		offset := (req.Current - 1) * req.PageSize
		err = t.db.Where(query, args...).Limit(req.PageSize, offset).Find(&list)
	} else {
		err = errors.New("is empty")
	}

	return
}
