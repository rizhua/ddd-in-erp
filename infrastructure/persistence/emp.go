package persistence

import (
	"errors"
	"strings"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

func NewEmp(db *xorm.Engine) repository.Emp {
	// db.SetSchema("public")
	return &emp{db: db}
}

type emp struct {
	db *xorm.Engine
}

func (t *emp) Create(info po.Emp) error {
	has, err := t.db.Insert(&info)
	if err != nil || has == 0 {
		return errors.New("写入失败")
	}

	return nil
}

func (t *emp) Update(info po.Emp) error {
	has, err := t.db.Where("id=? AND org_id=?", info.ID, info.OrgID).Omit("org_id").Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

func (t *emp) Get(info po.Emp) (po.Emp, error) {
	has, err := t.db.Get(&info)
	if err != nil || !has {
		err = errors.New("查询失败")
	}
	return info, err
}

func (t *emp) Find(req query.Request) (list []po.Emp, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "orgId":
			query += " AND a.org_id=?"
			args = append(args, v.Value)
		case "deptId":
			query += " AND b.dept_id=?"
			args = append(args, v.Value)
		case "keyword":
			query += " AND (a.name LIKE ? OR a.number LIKE ?)"
			args = append(args, v.Value.(string)+"%", v.Value.(string)+"%")
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	// 统计条数
	emp := new(po.Emp)
	total, err = t.db.Alias("a").Join("LEFT", "dept_emp AS b", "b.emp_id=a.id").Where(query, args...).Count(emp)
	if err != nil {
		return
	}
	if total > 0 {
		// 分页数据
		offset := (req.Current - 1) * req.PageSize
		err = t.db.Alias("a").Join("LEFT", "dept_emp AS b", "b.emp_id=a.id").Where(query, args...).Limit(req.PageSize, offset).Find(&list)
	} else {
		err = errors.New("数据为空")
	}

	return
}

func (t *emp) FindByUserID(userID int64) (list []po.Emp, err error) {
	err = t.db.Where("user_id=?", userID).Find(&list)
	return
}
