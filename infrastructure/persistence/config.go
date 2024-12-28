package persistence

import (
	"errors"
	"strings"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

func NewConfig(db *xorm.Engine) repository.Config {
	// db.SetSchema("public")
	return &config{db: db}
}

type config struct {
	db *xorm.Engine
}

// 创建配置
func (t *config) Create(info po.Config) error {
	has, err := t.db.Insert(&info)
	if err != nil || has == 0 {
		return errors.New("写入失败")
	}
	return nil
}

// 删除配置
func (t *config) Delete(id []int64) error {
	info := new(po.Config)
	has, err := t.db.In("id", id).Delete(info)
	if err != nil || has == 0 {
		return errors.New("删除失败")
	}

	return nil
}

// 更新配置
func (t *config) Update(info po.Config) error {
	has, err := t.db.ID(info.ID).Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

// 配置详情
func (t *config) Get(info po.Config) (po.Config, error) {
	has, err := t.db.Get(&info)
	if err != nil || !has {
		return info, errors.New("查询失败")
	}

	return info, nil
}

// 配置列表
func (t *config) Find(req query.Request) (list []po.Config, total int64, err error) {
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

	config := new(po.Config)

	// 统计条数
	total, err = t.db.Where(query, args...).Count(config)
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
