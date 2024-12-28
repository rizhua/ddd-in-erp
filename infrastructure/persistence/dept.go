package persistence

import (
	"errors"
	"strings"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

func NewDept(db *xorm.Engine) repository.Dept {
	// db.SetSchema("public")
	return &dept{db: db}
}

type dept struct {
	db *xorm.Engine
}

// 新增部门
func (t *dept) Create(info po.Dept) error {
	has, err := t.db.Insert(&info)
	if err != nil || has == 0 {
		return errors.New("写入失败")
	}

	return nil
}

// 删除部门
func (t *dept) Delete(orgID int64, id []int64) error {
	info := new(po.Dept)
	has, err := t.db.Where("org_id=?", orgID).In("id", id).Delete(info)
	if err != nil || has == 0 {
		return errors.New("删除失败")
	}

	return nil
}

// 更新部门
func (t *dept) Update(info po.Dept) error {
	has, err := t.db.Where("id=? AND org_id=?", info.ID, info.OrgID).Omit("org_id").Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

// 部门列表
func (t *dept) Find(orgID int64, req query.Request) (list []po.Dept, err error) {
	var (
		query string
		args  []any
	)

	query += "org_id=?"
	args = append(args, orgID)
	for _, v := range req.QueryBy {
		switch v.Field {
		case "name":
			query += " AND name LIKE ?"
			args = append(args, v.Value.(string)+"%")
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	err = t.db.Where(query, args...).Find(&list)
	if err != nil {
		err = errors.New("查询失败")
		return
	}

	return
}
