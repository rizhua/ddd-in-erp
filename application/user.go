package application

import (
	"context"
	"encoding/json"

	"rizhua.com/application/assembler"
	"rizhua.com/application/command"
	"rizhua.com/application/query"
	"rizhua.com/domain"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/converter"
)

func NewUserService(
	userDomain domain.UserService,
	userRepo repository.User,
) UserService {
	return UserService{
		userDomain: userDomain,
		userRepo:   userRepo,
	}
}

type UserService struct {
	Context    context.Context
	userDomain domain.UserService
	userRepo   repository.User
}

func (t *UserService) SignIn(args []byte) (any, error) {
	req := query.LoginUser{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return t.userDomain.SignIn(req)
}

func (t *UserService) SignUp(args []byte) (any, error) {
	cmd := command.CreateUser{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return nil, err
	}
	return t.userDomain.SignUp(cmd)
}

func (t *UserService) Active(args []byte) error {
	var token string
	return t.userDomain.Active(token)
}

func (t *UserService) Find(args []byte) (any, error) {
	req := query.Request{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	ret, cnt, err := t.userRepo.Find(req)
	if err != nil {
		return nil, err
	}
	data := make(map[string]any)
	data["list"] = new(converter.User).List(ret)
	data["total"] = cnt

	return data, err
}

// 忘记密码
func (t *UserService) Forget(args []byte) (string, error) {
	req := query.ForgetUser{}
	if err := json.Unmarshal(args, &req); err != nil {
		return "", err
	}
	if err := req.Validate(); err != nil {
		return "", err
	}
	return t.userDomain.Forget(req)
}

// 重置密码
func (t *UserService) RePassword(args []byte) error {
	cmd := command.ResetUserPassword{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	if err := cmd.Validate(); err != nil {
		return err
	}
	return t.userDomain.RePassword(cmd.Token, cmd.Password)
}

// 修改密码
func (t *UserService) SetPassword(args []byte) error {
	cmd := command.UpdateUserPassword{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	if err := cmd.Validate(); err != nil {
		return err
	}
	t.userDomain.Context = t.Context
	return t.userDomain.SetPassword(cmd.OldPassword, cmd.NewPassword)
}

// 供职组织
func (t *UserService) Work() (any, error) {
	t.userDomain.Context = t.Context
	ret, err := t.userDomain.Work()
	if err != nil {
		return nil, err
	}
	data := new(assembler.Org).List(ret)

	return data, nil
}

// 获取当前用户信息
func (t *UserService) Parse() (any, error) {
	t.userDomain.Context = t.Context
	ret, err := t.userDomain.Parse()
	if err != nil {
		return nil, err
	}
	return new(assembler.Login).Info(ret), nil
}
