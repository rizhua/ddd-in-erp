package persistence

import (
	"errors"
	"strings"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

func NewSpu(db *xorm.Engine) repository.Spu {
	return &spu{db: db}
}

type spu struct {
	db *xorm.Engine
}

func (t *spu) Create(info po.Spu) error {
	has, err := t.db.Insert(&info)
	if err != nil || has == 0 {
		return errors.New("写入失败")
	}
	return nil
}

func (t *spu) Delete(id []int64) error {
	spuPO := po.Spu{}
	has, err := t.db.In("id", id).Delete(&spuPO)
	if err != nil || has == 0 {
		return errors.New("删除失败")
	}
	return nil
}

func (t *spu) Update(info po.Spu) error {
	has, err := t.db.Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func (t *spu) Get(id int64) (info po.Spu, err error) {
	has, err := t.db.Where("id = ?", id).Get(&info)
	if err != nil || !has {
		err = errors.New("查询失败")
	}
	return
}

func (t *spu) Find(req query.Request) (list []po.Spu, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "title":
			query += " AND name LIKE ?"
			args = append(args, v.Value.(string)+"%")
		case "code":
			query += " AND code=?"
			args = append(args, v.Value, v.Value)
		case "barcode":
			query += " AND barcode=?"
			args = append(args, v.Value)
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	// 统计条数
	total, err = t.db.Where(query, args...).Count(new(po.Spu))
	if err != nil {
		return
	}
	if total > 0 {
		// 分页数据
		offset := (req.Current - 1) * req.PageSize
		err = t.db.Where(query, args...).Omit("attribute,detail").Limit(req.PageSize, offset).Find(&list)
	} else {
		err = errors.New("is empty")
	}

	return
}
