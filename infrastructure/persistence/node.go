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

func NewNode(db *xorm.Engine) repository.Node {
	// db.SetSchema("public")
	return &node{db: db}
}

type node struct {
	db *xorm.Engine
}

// 创建节点
func (t *node) Create(info po.Node) error {
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
		node := new(po.Node)
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

// 删除节点
func (t *node) Delete(id []int64) error {
	has, err := t.db.In("id", id).Delete(new(po.Node))
	if err != nil || has == 0 {
		return errors.New("删除失败")
	}

	return nil
}

// 更新节点
func (t *node) Update(info po.Node) error {
	has, err := t.db.ID(info.ID).Omit("parent_id").Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

// 节点列表
func (t *node) Find(req query.Request) (list []po.Node, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "name":
			query += " AND a.name LIKE ?"
			args = append(args, v.Value.(string)+"%")
		case "meta":
			query += " AND a.meta=?"
			args = append(args, v.Value)
		case "path":
			query += " AND (a.path=? OR a.path LIKE ?)"
			args = append(args, v.Value, "%,"+v.Value.(string))
		case "parentID":
			query += " AND a.parent_id=?"
			args = append(args, v.Value)
		case "bundleId":
			err = t.db.Alias("a").Join("LEFT", "bundle_node as b", "b.node_id=a.id").Where("b.bundle_id=?", v.Value).OrderBy("sort").Find(&list)
			return
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	// 统计条数
	total, err = t.db.Alias("a").Where(query, args...).Count(new(po.Node))
	if err != nil {
		return
	}
	if total > 0 {
		// 分页数据
		offset := (req.Current - 1) * req.PageSize
		err = t.db.Alias("a").Where(query, args...).OrderBy("sort").Limit(req.PageSize, offset).Find(&list)
	} else {
		err = errors.New("is empty")
	}

	return
}

// 节点排序
func (t *node) SetSort(info po.Node) error {
	has, err := t.db.ID(info.ID).Cols("sort").Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

// 修改节点状态
func (t *node) SetStatus(info po.Node) error {
	has, err := t.db.ID(info.ID).Cols("status").Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

// 通过scheme取得应用
func (t *node) GetByMeta(meta string) (info po.Node, err error) {
	where := "meta=?"
	has, err := t.db.Where(where, meta).Get(&info)
	if err != nil || !has {
		err = errors.New("查询失败")
	}

	return
}

// 权限
func (t *node) Permission(userID int64, path string) (list []po.Node, err error) {
	if userID == 0 {
		err = t.db.Where("(path=? OR path LIKE ?) AND status>0", path, "%,"+path).OrderBy("sort").Find(&list)
	} else {
		db := t.db.Alias("a").Join("LEFT", "role_node AS b", "b.node_id=a.id").Join("LEFT", "role_user AS c", "c.role_id=b.role_id")
		err = db.Where("c.user_id=? AND (a.path=? OR a.path LIKE ?) AND status>0", userID, path, "%,"+path).OrderBy("a.sort").Find(&list)
	}
	if err != nil {
		err = errors.New("查询失败")
		return
	}
	if len(list) == 0 {
		err = errors.New("数据为空")
	}

	return
}
