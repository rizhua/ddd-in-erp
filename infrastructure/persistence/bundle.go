package persistence

import (
	"errors"
	"strings"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

func NewBundle(db *xorm.Engine) repository.Bundle {
	// db.SetSchema("public")
	return &bundle{db: db}
}

type bundle struct {
	db *xorm.Engine
}

// 新增套餐
func (t *bundle) Create(info po.Bundle) error {
	row, err := t.db.Insert(&info)
	if nil != err || row == 0 {
		err = errors.New("写入失败")
	}
	return err
}

// 删除套餐
func (t *bundle) Delete(id []int64) error {
	info := new(po.Bundle)
	has, err := t.db.In("id", id).Delete(info)
	if err != nil || has == 0 {
		return errors.New("删除失败")
	}

	return nil
}

// 更新套餐
func (t *bundle) Update(info po.Bundle) error {
	row, err := t.db.Where("id=?", info.ID).Update(&info)
	if err != nil || row == 0 {
		err = errors.New("更新失败")
	}
	return err
}

// 套餐信息
func (t *bundle) Get(req query.Request) (info po.Bundle, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "id":
			query += " AND id=?"
			args = append(args, v.Value)
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	has, err := t.db.Where(query, args...).Get(&info)
	if err != nil || !has {
		err = errors.New("数据为空")
		return
	}

	return
}

// 套餐列表
func (t *bundle) Find(req query.Request) (list []po.Bundle, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "name":
			query += " AND a.name LIKE ?"
			args = append(args, v.Value.(string)+"%")
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	// 统计条数
	total, err = t.db.Alias("a").Where(query, args...).Count(new(po.Bundle))
	if err != nil {
		return
	}
	if total > 0 {
		// 分页数据
		offset := (req.Current - 1) * req.PageSize
		err = t.db.Alias("a").Select("a.*,(select count(id) from license as b where b.biz_id=a.id) as license_count").Where(query, args...).Limit(req.PageSize, offset).Find(&list)
	} else {
		err = errors.New("数据为空")
	}

	return
}

// 某套餐nodeID
func (t *bundle) FindNodeID(bundleID int64) (list []po.BundleNode, err error) {
	err = t.db.Where("bundle_id=?", bundleID).Find(&list)
	if err != nil {
		err = errors.New("查询失败")
	}
	if len(list) == 0 {
		err = errors.New("数据为空")
	}
	return
}

// 绑定、解绑节点
func (t *bundle) BindNodeID(bundleID int64, nodeID []int64) error {
	session := t.db.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	info := new(po.BundleNode)
	// 1、 清空
	_, err := session.Where("bundle_id=?", bundleID).Delete(info)
	if err != nil {
		session.Rollback()
		return errors.New("清空失败")
	}
	// 2、 写入
	list := make([]*po.BundleNode, len(nodeID))
	for i, id := range nodeID {
		list[i] = &po.BundleNode{BundleID: bundleID, NodeID: id}
	}
	has, err := session.Insert(list)
	if err != nil || has == 0 {
		session.Rollback()
		return errors.New("写入失败")
	}
	return session.Commit()
}

// 某套餐节点列表
func (t *bundle) FindNode(bundleID int64, path string) (list []po.Node, err error) {
	var (
		query string
		args  []any
	)

	query = "b.bundle_id=?"
	args = append(args, bundleID)
	if path != "" {
		query += " AND (a.path=? OR a.path LIKE ?)"
		args = append(args, path, "%,"+path)
	}
	err = t.db.Alias("a").Join("LEFT", "bundle_node AS b", "b.node_id=a.id").Where(query, args...).Find(&list)
	if err != nil {
		err = errors.New("查询失败")
	}
	if len(list) == 0 {
		err = errors.New("数据为空")
	}
	return
}

// 许可证信息
func (t *bundle) GetLicense(req query.Request) (info po.License, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "name":
			query += " AND a.name LIKE ?"
			args = append(args, v.Value.(string)+"%")
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	has, err := t.db.Where(query, args...).Get(&info)
	if err != nil || !has {
		err = errors.New("查询失败")
	}
	return
}

// 许可证列表
func (t *bundle) FindLicense(req query.Request) (list []po.License, cnt int64, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "code":
			query += " AND code LIKE ?"
			args = append(args, v.Value.(string)+"%")
		case "bizId":
			query += " AND biz_id=?"
			args = append(args, v.Value)
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	// 统计条数
	cnt, err = t.db.Where(query, args...).Count(new(po.License))
	if err != nil {
		return
	}
	if cnt > 0 {
		// 分页数据
		offset := (req.Current - 1) * req.PageSize
		err = t.db.Where(query, args...).Limit(req.PageSize, offset).Find(&list)
	} else {
		err = errors.New("数据为空")
	}

	return
}
