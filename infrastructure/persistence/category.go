package persistence

import (
	"errors"
	"fmt"
	"strings"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

func NewCategory(db *xorm.Engine) repository.Category {
	return &category{db: db}
}

type category struct {
	db *xorm.Engine
}

// 创建类目
func (t *category) Create(info po.Category) error {
	session := t.db.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}

	has, err := session.Insert(&info)
	if err != nil || has == 0 {
		return errors.New("写入失败")
	}

	if info.ParentID == 0 {
		info.Path = fmt.Sprintf("%d", info.ID)
	} else {
		node := new(po.Category)
		has, err := session.Where("id=?", info.ParentID).Get(node)
		if err != nil || !has {
			session.Rollback()
			return errors.New("父级不存在")
		}
		info.Path = fmt.Sprintf("%d,%s", info.ID, node.Path)
	}
	has, err = session.ID(info.ID).Cols("path").Update(&info)
	if err != nil || has == 0 {
		session.Rollback()
		return errors.New("写入失败")
	}

	return session.Commit()
}

// 删除类目
func (t *category) Delete(id []int64) error {
	info := new(po.Category)
	has, err := t.db.In("id", id).Delete(info)
	if err != nil || has == 0 {
		return errors.New("删除失败")
	}

	return nil
}

// 更新类目
func (t *category) Update(info po.Category) error {
	has, err := t.db.Where("id=?", info.ID).Omit("parent_id").Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

// 更新类目排序
func (t *category) SetSort(id int64, sort int8) error {
	info := po.Category{
		Sort: sort,
	}
	has, err := t.db.Where("id=?", id).Cols("sort").Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

// 类目列表
func (t *category) Find(req query.Request) (list []po.Category, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "name":
			query += "name LIKE ?"
			args = append(args, v.Value.(string)+"%")
		case "parentId":
			query += " AND parent_id=?"
			args = append(args, v.Value)
		case "orgId":
			query += " AND org_id=?"
			args = append(args, v.Value)
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	info := new(po.Category)

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
		err = errors.New("数据为空")
	}

	return
}

// 新增类目属性
func (t *category) CreateAttribute(info po.CategoryAttribute) error {
	has, err := t.db.Insert(&info)
	if err != nil || has == 0 {
		return errors.New("写入失败")
	}

	return nil
}

// 更新类目属性
func (t *category) UpdateAttribute(info po.CategoryAttribute) error {
	has, err := t.db.Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

// 删除类目属性
func (t *category) DeleteAttribute(categoryID int64, id []int64) error {
	info := new(po.CategoryAttribute)
	has, err := t.db.In("id", id).Where("category_id=?", categoryID).Delete(info)
	if err != nil || has == 0 {
		return errors.New("删除失败")
	}

	return nil
}

// 类目属性列表
func (t *category) FindAttribute(req query.Request) (list []po.CategoryAttribute, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "categoryId":
			query += "category_id=?"
			args = append(args, v.Value)
		case "label":
			query += "label LIKE ?"
			args = append(args, v.Value.(string)+"%")
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	info := new(po.CategoryAttribute)

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
		err = errors.New("数据为空")
	}

	return
}
