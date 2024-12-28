package persistence

import (
	"errors"
	"strings"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

type user struct {
	db *xorm.Engine
}

func NewUser(db *xorm.Engine) repository.User {
	// db.SetSchema("public")
	return &user{db: db}
}

func (t *user) Create(info po.User) (id int64, err error) {
	has, err := t.db.Insert(&info)
	if err != nil || has == 0 {
		err = errors.New("写入失败")
		return
	}
	id = info.ID
	return
}

func (t *user) Update(info po.User) error {
	has, err := t.db.ID(info.ID).Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

func (t *user) Get(req po.User) (info po.User, err error) {
	var (
		query string
		args  []interface{}
	)

	if req.ID > 0 {
		query += " AND id=?"
		args = append(args, req.ID)
	}
	if req.Email != "" {
		query += " AND email=?"
		args = append(args, req.Email)
	}
	if req.Mobile != "" {
		query += " AND mobile=?"
		args = append(args, req.Mobile)
	}
	if req.Password != "" {
		query += " AND password=?"
		args = append(args, req.Password)
	}
	query = strings.TrimPrefix(query, " AND ")

	has, err := t.db.Where(query, args...).Get(&info)
	if err != nil {
		err = errors.New("查询失败")
	}

	if !has {
		err = errors.New("数据为空")
	}

	return
}

// 根据账号获取用户ID
func (t *user) GetByAccount(account string) (int64, error) {
	info := new(po.User)
	has, err := t.db.Where("mobile=? OR email=?", account, account).Get(info)
	if err != nil || !has {
		err = errors.New("查询失败")
	}

	return info.ID, err
}

func (t *user) Find(req query.Request) (list []po.User, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "nickname":
			query += " AND nickname LIKE ?"
			args = append(args, v.Value.(string)+"%")
		case "mobile":
			query += " AND mobile LIKE ?"
			args = append(args, v.Value.(string)+"%")
		case "email":
			query += " AND email LIKE ?"
			args = append(args, v.Value.(string)+"%")
		}
	}
	query = strings.TrimPrefix(query, " AND ")
	user := new(po.User)

	// 统计条数
	total, err = t.db.Where(query, args...).Count(user)
	if err != nil {
		return
	}
	if total > 0 {
		// 分页数据
		offset := (req.Current - 1) * req.PageSize
		err = t.db.Omit("password", "access_key", "secret_key").Where(query, args...).Limit(req.PageSize, offset).Find(&list)
	} else {
		err = errors.New("is empty")
	}

	return
}

func (t *user) SetPassword(info po.User) error {
	has, err := t.db.ID(info.ID).Cols("password").Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

func (t *user) Exist(info po.User) bool {
	has, err := t.db.Exist(&info)
	if err != nil || !has {
		return false
	}
	return true
}
