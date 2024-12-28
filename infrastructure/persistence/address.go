package persistence

import (
	"errors"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

type address struct {
	db *xorm.Engine
}

func NewAddress(db *xorm.Engine) repository.Address {
	// db.SetSchema("public")
	return &address{db: db}
}

// 新增收、发地址
func (t *address) Create(info po.Address) error {
	row, err := t.db.Insert(&info)
	if err != nil || row == 0 {
		err = errors.New("写入失败")
	}
	return err
}

// 删除收、发地址
func (t *address) Delete(id []int64, userID, orgID int64) error {
	address := new(po.Address)
	row, err := t.db.Where("user_id=? OR org_id=?", userID, orgID).In("id", id).Delete(address)
	if err != nil || row == 0 {
		err = errors.New("删除失败")
	}
	return err
}

// 更新收、发地址
func (t *address) Update(info po.Address) error {
	row, err := t.db.Omit("default").Where("id=? AND user_id=?", info.ID, info.UserID).Update(&info)
	if nil != err || row == 0 {
		err = errors.New("更新失败")
	}
	return err
}

// 设置默认收、发地址
func (t *address) SetDefault(id, orgID, userID int64) error {
	session := t.db.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}

	info := po.Address{Default: false}
	row, err := session.Cols("default").Where("org_id=? OR user_id=?", orgID, userID).Update(&info)
	if err != nil || row == 0 {
		session.Rollback()
		return errors.New("更新失败")
	}

	info.Default = true
	_, err = session.Cols("default").Where("id= AND (org_id=? OR user_id=?)", info.ID, orgID, userID).NoAutoTime().Update(&info)
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

// 收、发地址列表
func (t *address) Find(userID, orgID int64, req query.Request) (list []po.Address, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	query = "(user_id=? OR org_id=?)"
	args = append(args, userID, orgID)
	for _, v := range req.QueryBy {
		switch v.Field {
		case "contact":
			query += " AND contact LIKE ?"
			args = append(args, v.Value.(string)+"%")
		case "tel":
			query += " AND tel LIKE ?"
			args = append(args, v.Value.(string)+"%")
		}
	}

	// 统计条数
	address := new(po.Address)
	total, err = t.db.Where(query, args...).Count(address)
	if err != nil {
		return
	}
	if total > 0 {
		// 分页数据
		offset := (req.Current - 1) * req.PageSize
		err = t.db.Where(query, args...).Limit(req.PageSize, offset).Desc("id").Find(&list)
	} else {
		err = errors.New("数据为空")
	}

	return
}
