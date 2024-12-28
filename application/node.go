package application

import (
	"context"
	"encoding/json"

	"rizhua.com/application/command"
	"rizhua.com/application/query"
	"rizhua.com/domain"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/converter"
	"rizhua.com/infrastructure/persistence/po"
)

func NewNodeService(
	nodeDomain domain.NodeService,
	nodeRepo repository.Node,
	userDomain domain.UserService,
) NodeService {
	return NodeService{
		nodeDomain: nodeDomain,
		nodeRepo:   nodeRepo,
		userDomain: userDomain,
	}
}

type NodeService struct {
	Context    context.Context
	nodeDomain domain.NodeService
	nodeRepo   repository.Node
	userDomain domain.UserService
}

func (t *NodeService) Create(args []byte) error {
	cmd := command.CreateNode{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	return t.nodeDomain.Create(cmd)
}

func (t *NodeService) Delete(args []byte) error {
	cmd := command.Delete{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	if err := cmd.Validate(); err != nil {
		return err
	}
	return t.nodeRepo.Delete(cmd.ID)
}

func (t *NodeService) Update(args []byte) error {
	cmd := command.UpdateNode{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	return t.nodeDomain.Update(cmd)
}

func (t *NodeService) Find(args []byte) (any, error) {
	req := query.Request{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	data := make(map[string]any)
	ret, cnt, err := t.nodeRepo.Find(req)
	if err != nil {
		return data, err
	}

	data["list"] = new(converter.Node).Tree(ret, 0)
	data["total"] = cnt

	return data, err
}

func (t *NodeService) SetSort(args []byte) error {
	cmd := command.UpdateNodeSort{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	nodePO := po.Node{ID: cmd.ID, Sort: cmd.Sort}
	return t.nodeRepo.SetSort(nodePO)
}

func (t *NodeService) SetStatus(args []byte) error {
	cmd := command.UpdateNodeStatus{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	nodePO := po.Node{ID: cmd.ID, Status: cmd.Status}
	return t.nodeRepo.SetStatus(nodePO)
}

func (t *NodeService) Permission(args []byte) (data any, err error) {
	req := query.Permission{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	node, err := t.nodeRepo.GetByMeta(req.Meta)
	if err != nil {
		return nil, err
	}
	t.nodeDomain.Context = t.Context
	ret, err := t.nodeDomain.Permission(node.Path)
	if err != nil {
		return nil, err
	}
	tmp := new(converter.Node).Tree(ret, 0)
	if len(tmp) > 0 {
		data = tmp[0].Children
	}

	return
}
