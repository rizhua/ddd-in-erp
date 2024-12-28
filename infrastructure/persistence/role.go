package persistence

import (
	"errors"
	"strings"

	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

type role struct {
	db *xorm.Engine
}

func NewRole(db *xorm.Engine) repository.Role {
	// db.SetSchema("public")
	return &role{db: db}
}

// 新增角色
func (t *role) Create(info po.Role) error {
	has, err := t.db.Insert(&info)
	if err != nil || has == 0 {
		return errors.New("写入失败")
	}

	return nil
}

// 删除角色
func (t *role) Delete(orgID int64, id []int64) error {
	role := new(po.Role)
	has, err := t.db.Where("org_id=?", orgID).In("id", id).Delete(role)
	if err != nil || has == 0 {
		return errors.New("删除失败")
	}

	return nil
}

// 更新角色
func (t *role) Update(info po.Role) error {
	has, err := t.db.Where("id=? AND org_id=?", info.ID, info.OrgID).Omit("org_id").Update(&info)
	if err != nil || has == 0 {
		return errors.New("更新失败")
	}

	return nil
}

// 角色详情
func (t *role) GetByID(id int64) (info po.Role, err error) {
	has, err := t.db.ID(id).Get(&info)
	if err != nil || !has {
		err = errors.New("查询失败")
		return
	}

	return
}

// 角色列表
func (t *role) Find(req query.Request) (list []po.Role, err error) {
	var (
		query string
		args  []interface{}
	)

	for _, v := range req.QueryBy {
		switch v.Field {
		case "orgId":
			query += "org_id=?"
			args = append(args, v.Value)
		case "name":
			query += " AND name LIKE ?"
			args = append(args, v.Value.(string)+"%")
		case "parentID":
			query += " AND parent_id=?"
			args = append(args, v.Value)
		}
	}
	query = strings.TrimPrefix(query, " AND ")

	err = t.db.Where(query, args...).Or("org_id=?", 0).Find(&list)
	if err != nil {
		err = errors.New("查询失败")
		return
	}

	return
}

// 绑定、解绑节点
func (t *role) BindNodeID(roleID int64, nodeID []int64) error {
	session := t.db.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	info := new(po.RoleNode)
	_, err := session.Where("role_id=?", roleID).Delete(info)
	if err != nil {
		session.Rollback()
		return errors.New("删除失败")
	}
	infos := make([]*po.RoleNode, len(nodeID))
	for i, id := range nodeID {
		infos[i] = &po.RoleNode{RoleID: roleID, NodeID: id}
	}
	has, err := session.Insert(infos)
	if err != nil || has == 0 {
		session.Rollback()
		return errors.New("写入失败")
	}

	return session.Commit()
}

// 节点列表
func (t *role) FindNode(roleID int64, req query.Request) (list []po.Node, err error) {
	var (
		query string
		args  []interface{}
	)

	query = "b.role_id=?"
	args = append(args, roleID)
	for _, v := range req.QueryBy {
		switch v.Field {
		case "name":
			query += " AND a.name LIKE ?"
			args = append(args, v.Value.(string)+"%")
		}
	}

	err = t.db.Alias("a").Where(query, args...).Join("LEFT", "role_node AS b", "b.node_id=a.id").Find(&list)
	if len(list) == 0 {
		err = errors.New("数据为空")
	}

	return
}

// 节点 id 列表
func (t *role) FindNodeID(roleID int64) (list []po.RoleNode, err error) {
	err = t.db.Where("role_id=?", roleID).Find(&list)
	if len(list) == 0 {
		err = errors.New("数据为空")
	}

	return
}

// 添加成员
func (t *role) AddUser(roleID int64, userID []int64) error {
	roleUsers := make([]po.RoleUser, len(userID))
	for i, v := range userID {
		roleUsers[i] = po.RoleUser{
			RoleID: roleID,
			UserID: v,
		}
	}
	has, err := t.db.Insert(&roleUsers)
	if err != nil || has == 0 {
		return errors.New("写入失败")
	}

	return nil
}

// 移除成员
func (t *role) RemoveUser(roleID int64, nodeID []int64) error {
	info := po.RoleUser{RoleID: roleID}
	has, err := t.db.In("user_id", nodeID).Delete(&info)
	if err != nil || has == 0 {
		return errors.New("删除失败")
	}

	return nil
}

// 成员列表
func (t *role) FindUser(orgID, roleID int64, req query.Request) (list []po.Emp, total int64, err error) {
	var (
		query string
		args  []interface{}
	)

	query = "a.org_id=? AND b.role_id=?"
	args = append(args, orgID, roleID)
	for _, v := range req.QueryBy {
		switch v.Field {
		case "number":
			query += " AND a.number LIKE ?"
			args = append(args, v.Value.(string)+"%")
		case "mobile":
			query += " AND a.mobile=?"
			args = append(args, v.Value)
		}
	}

	// 统计条数
	emp := new(po.Emp)
	total, err = t.db.Alias("a").Join("INNER", "role_user AS b", "b.user_id=a.user_id").Where(query, args...).Count(emp)

	if err != nil {
		return
	}

	if total > 0 {
		// 分页数据
		offset := (req.Current - 1) * req.PageSize
		err = t.db.Alias("a").Join("INNER", "role_user AS b", "b.user_id=a.user_id").Where(query, args...).Limit(req.PageSize, offset).Find(&list)
	} else {
		err = errors.New("数据为空")
	}

	return
}
