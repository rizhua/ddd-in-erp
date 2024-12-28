package persistence

import (
	"errors"
	"strings"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

func NewBrand(db *xorm.Engine) repository.Brand {
	return &brand{db: db}
}

type brand struct {
	db *xorm.Engine
}

// 创建品牌
func (t *brand) Create(info po.Brand) error {
	has, err := t.db.Insert(&info)
	if err != nil || has == 0 {
		return errors.New("写入失败")
	}

	return err
}

// 删除品牌
func (t *brand) Delete(id []int64) error {
	brand := new(po.Brand)
	has, err := t.db.In("id", id).Delete(brand)
	if err != nil || has == 0 {
		return errors.New("删除失败")
	}

	return nil
}

// 更新品牌
func (t *brand) Update(info po.Brand) error {
	has, err := t.db.ID(info.ID).Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

// 品牌列表
func (t *brand) Find(req query.Request) (list []po.Brand, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "name":
			query += " AND name LIKE ?"
			args = append(args, v.Value.(string)+"%")
		case "letter":
			query += " AND letter=?"
			args = append(args, v.Value)
		case "status":
			query += " AND status=?"
			args = append(args, v.Value)
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	brand := new(po.Brand)
	// 统计条数
	total, err = t.db.Where(query, args...).Count(brand)
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
