package persistence

import (
	"errors"
	"strings"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

func NewOrg(db *xorm.Engine) repository.Org {
	// db.SetSchema("public")
	return &org{db: db}
}

type org struct {
	db *xorm.Engine
}

// 创建组织
func (t *org) Create(info po.Org) error {
	has, err := t.db.Insert(&info)
	if err != nil || has == 0 {
		return errors.New("写入失败")
	}
	return nil
}

// 更新组织
func (t *org) Update(info po.Org) error {
	has, err := t.db.ID(info.ID).Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

func (t *org) GetByID(id int64) (info po.Org, err error) {
	has, err := t.db.Where("id=?", id).Get(&info)
	if err != nil || !has {
		err = errors.New("查询失败")
	}

	return
}

func (t *org) GetByUserID(id, userID int64) (info po.Org, err error) {
	has, err := t.db.Alias("a").Join("LEFT", "emp AS b", "b.org_id=a.id").Where("a.id=? AND b.user_id=?", id, userID).Get(&info)
	if err != nil || !has {
		err = errors.New("查询失败")
	}

	return
}

// 组织列表
func (t *org) Find(req query.Request) (list []po.Org, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "code":
			query += " AND name LIKE ?"
			args = append(args, v.Value.(string)+"%")
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	info := new(po.Org)

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

func (t *org) Count(req query.Request) (int64, error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "code":
			query += " AND name LIKE ?"
			args = append(args, v.Value.(string)+"%")
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	return t.db.Where(query, args...).Count(new(po.Org))
}

func (t *org) FindByUserID(userID int64) (list []po.Org, err error) {
	err = t.db.Alias("a").Join("LEFT", "emp as b", "b.org_id=a.id").Where("b.user_id=?", userID).Find(&list)
	if len(list) == 0 {
		err = errors.New("数据为空")
	}
	return
}
