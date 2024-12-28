package domain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"rizhua.com/application/command"
	"rizhua.com/application/query"
	"rizhua.com/domain/entity"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/adapter"
	"rizhua.com/infrastructure/constant"
	"rizhua.com/infrastructure/etc"
	"rizhua.com/pkg/util"
)

func NewUserService(
	userRepo repository.User,
	empRepo repository.Emp,
	orgRepo repository.Org,
) UserService {
	entity.UserRepo = userRepo
	entity.EmpRepo = empRepo
	entity.OrgRepo = orgRepo
	return UserService{
		userRepo: userRepo,
		empRepo:  empRepo,
		orgRepo:  orgRepo,
	}
}

type UserService struct {
	userRepo repository.User
	empRepo  repository.Emp
	orgRepo  repository.Org
	Context  context.Context
}

func (t *UserService) SignIn(req query.LoginUser) (token string, err error) {
	userDO := entity.User{}
	if req.Account == "root" {
		cfg := etc.C
		if cfg.Root.Password != req.Password {
			err = errors.New("密码不正确")
			return
		}
		userDO.Nickname = "root"
	} else {
		if util.IsMobile(req.Account) {
			userDO.Mobile = req.Account
		}
		if util.IsEmail(req.Account) {
			userDO.Email = req.Account
		}
		// 0-密码登录，1-验证码登录，2-微信登录，3-扫码登录
		switch req.Mode {
		case 1:
			err = userDO.WithCaptcha(req.Captcha)
		case 2:
			err = errors.New("暂不支持微信登录")
		case 3:
			err = errors.New("暂不支持扫码登录")
		default:
			userDO.Password = req.Password
			err = userDO.WithPassword()
		}
		if err != nil {
			return
		}
		userDO, err = userDO.Info()
		if err != nil {
			return
		}
		if userDO.Status == 0 {
			err = errors.New("账号已被禁用")
			return
		}
	}
	// 缓存登录信息
	cache := adapter.NewCache()
	key := fmt.Sprintf("%v", constant.USER)
	field := fmt.Sprintf("%d", userDO.ID)
	value, _ := json.Marshal(userDO)
	err = cache.HSet(key, field, value)
	if err != nil {
		return
	}

	return userDO.MakeToken()
}

func (t *UserService) SignUp(cmd command.CreateUser) (token string, err error) {
	userDO := entity.User{}
	// 邮箱注册
	if util.IsEmail(cmd.Account) {
		// userDO.Email = req.Account
		err = userDO.WithEmail()
		return
	}

	// 手机号注册
	err = userDO.WithMobile(cmd.Captcha, cmd.Account, cmd.Password)
	return
}

func (t *UserService) Active(token string) error {
	userDO := entity.User{}
	return userDO.Active(token)
}

func (t *UserService) Forget(req query.ForgetUser) (token string, err error) {
	userDO := entity.User{}
	if req.Email != "" {
		token, err = userDO.ForgetByEmail()
		// todo 发送邮件
		return
	}

	if req.Mobile != "" {
		return userDO.ForgetByMobile(req.Captcha, req.Mobile)
	}

	err = errors.New("账号不存在")
	return
}

func (t *UserService) RePassword(token, password string) error {
	var userDO = entity.User{
		Password: password,
	}
	// 校验token
	cache := adapter.NewCache()
	key := fmt.Sprintf("%s:%s", constant.TOKEN, token)
	if !cache.Exists(key) {
		return errors.New("无效 token")
	}
	userID, _ := cache.Get(key)
	userDO.ID, _ = strconv.ParseInt(userID, 10, 64)

	return userDO.SetPassword()
}

func (t *UserService) SetPassword(oldPassword, newPassword string) error {
	userDO, err := t.Parse()
	if err != nil {
		return err
	}

	// 校验旧密码
	if userDO.Password != oldPassword {
		return errors.New("旧密码不正确")
	}

	userDO.Password = newPassword
	err = userDO.SetPassword()
	return err
}

func (t *UserService) Work() (list []entity.Org, err error) {
	user, err := t.Parse()
	if err != nil {
		return
	}
	// 根据 userID 查询组织
	orgDO := entity.Org{}
	list, err = orgDO.FindByUserID(user.ID)
	if err != nil {
		return
	}
	if user.Org != nil {
		for _, v := range list {
			if v.ID == user.Org.ID {
				user.Org = &v
				break
			}
		}
	} else {
		user.Org = &list[0]
	}
	// 根据 userID,orgID 查询员工
	emp := entity.Emp{}
	user.Emp, err = emp.Get(user.Org.ID, user.ID)
	if err != nil {
		return
	}
	// 缓存登录信息
	cache := adapter.NewCache()
	key := fmt.Sprintf("%v", constant.USER)
	field := fmt.Sprintf("%d", user.ID)
	value, _ := json.Marshal(user)
	err = cache.HSet(key, field, value)

	return
}

// 解析登录信息
func (t *UserService) Parse() (entity.User, error) {
	userDO := entity.User{}
	return userDO.Parse(t.Context)
}
