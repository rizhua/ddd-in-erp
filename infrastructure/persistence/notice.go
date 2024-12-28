package persistence

import (
	"errors"
	"strings"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

func NewNotice(db *xorm.Engine) repository.Notice {
	// db.SetSchema("public")
	return &notice{db: db}
}

type notice struct {
	db *xorm.Engine
}

// 创建公告
func (t *notice) Create(info po.Notice) error {
	has, err := t.db.Insert(&info)
	if err != nil || has == 0 {
		return errors.New("写入失败")
	}
	return nil
}

// 删除公告
func (t *notice) Delete(id []int64) error {
	info := new(po.Notice)
	has, err := t.db.In("id", id).Delete(info)
	if err != nil || has == 0 {
		return errors.New("删除失败")
	}

	return nil
}

// 更新公告
func (t *notice) Update(info po.Notice) error {
	has, err := t.db.ID(info.ID).Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

// 公告列表
func (t *notice) Find(req query.Request) (list []po.Notice, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "title":
			query += " AND name LIKE ?"
			args = append(args, v.Value.(string)+"%")
		case "scope":
			query += " AND scope=?"
			args = append(args, v.Value, v.Value)
		case "type":
			query += " AND type=?"
			args = append(args, v.Value)
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	// 统计条数
	total, err = t.db.Where(query, args...).Count(new(po.Notice))
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
