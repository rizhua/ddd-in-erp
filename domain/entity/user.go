package entity

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/adapter"
	"rizhua.com/infrastructure/constant"
	"rizhua.com/infrastructure/persistence/po"
)

var UserRepo repository.User

// 根实体：用户
type User struct {
	ID        int64
	Nickname  string
	Mobile    string
	Email     string
	Password  string
	Birthday  time.Time
	Gender    int8
	AccessKey string
	SecretKey string
	Status    int8
	Org       *Org
	Emp       *Emp
}

func (t *User) Info() (User, error) {
	userPO := po.User{
		ID:       t.ID,
		Nickname: t.Nickname,
		Mobile:   t.Mobile,
		Email:    t.Email,
		Password: t.Password,
	}
	ret, err := UserRepo.Get(userPO)
	return User{
		ID:        ret.ID,
		Nickname:  ret.Nickname,
		Mobile:    ret.Mobile,
		Email:     ret.Email,
		Password:  ret.Password,
		Birthday:  ret.Birthday,
		Gender:    ret.Gender,
		AccessKey: ret.AccessKey,
		SecretKey: ret.SecretKey,
		Status:    ret.Status,
	}, err
}

func (t *User) SetPassword() (err error) {
	userPO := po.User{
		ID:       t.ID,
		Password: t.Password,
	}
	err = UserRepo.SetPassword(userPO)
	return
}

// 密码登录
func (t *User) WithPassword() error {
	info, err := t.Info()
	if err != nil {
		err = errors.New("账号或密码不正确")
	}
	t.ID = info.ID
	return err
}

// 验证码登录
func (t *User) WithCaptcha(code string) (err error) {
	cache := adapter.NewCache()
	key := fmt.Sprintf("%s:%s", constant.CAPTCHA, t.Mobile)
	captcha, _ := cache.Get(key)
	if captcha == "" || code != captcha {
		err = errors.New("验证码不正确")
		return
	}
	userPO := po.User{Mobile: t.Mobile}
	has := UserRepo.Exist(userPO)
	if !has {
		// 不存在即注册
		t.ID, err = UserRepo.Create(userPO)
		if err != nil {
			err = errors.New("注册失败")
			return
		}
	}
	cache.Del(key)
	return
}

// 通过邮箱找回密码
func (t *User) ForgetByEmail() (token string, err error) {
	userPO := po.User{Email: t.Email}
	info, err := UserRepo.Get(userPO)
	if err != nil || info.ID == 0 {
		err = errors.New("邮箱不存在")
		return
	}
	t.ID = info.ID
	return t.MakeToken()
}

// 通过手机号找回密码
func (t *User) ForgetByMobile(captcha, mobile string) (token string, err error) {
	userPO := po.User{Mobile: mobile}
	info, err := UserRepo.Get(userPO)
	if err != nil || info.ID == 0 {
		err = errors.New("手机号不存在")
		return
	}
	cache := adapter.NewCache()
	key := fmt.Sprintf("%s:%s", constant.CAPTCHA, mobile)
	rCaptcha, _ := cache.Get(key)
	if rCaptcha != captcha {
		err = errors.New("验证码错误")
		return
	}
	t.ID = info.ID
	return t.MakeToken()
}

func (t *User) MakeToken() (token string, err error) {
	token = uuid.NewString()
	cache := adapter.NewCache()
	key := fmt.Sprintf("%s:%s", constant.TOKEN, token)
	err = cache.Set(key, t.ID, 30*time.Minute)
	return
}

// 使用手机号注册
func (t *User) WithMobile(code, mobile, password string) error {
	cache := adapter.NewCache()
	key := fmt.Sprintf("%s:%s", constant.CAPTCHA, mobile)
	captcha, _ := cache.Get(key)
	if captcha == "" || code != captcha {
		return errors.New("验证码不正确")
	}
	userPO := po.User{
		Mobile: mobile,
	}
	has := UserRepo.Exist(userPO)
	if has {
		return errors.New("手机号已存在")
	}
	userPO.Password = password
	userID, err := UserRepo.Create(userPO)
	t.ID = userID
	return err
}

// 使用邮箱注册
func (t *User) WithEmail() error {
	return nil
}

func (t *User) Active(token string) error {
	cache := adapter.NewCache()
	key := fmt.Sprintf("%s:%s", constant.TOKEN, token)
	if !cache.Exists(key) {
		return errors.New("无效 token")
	}
	userPO := po.User{}
	userPO.Email, _ = cache.Get(key)
	has := UserRepo.Exist(userPO)
	if has {
		return errors.New("邮箱已存在")
	}
	userID, err := UserRepo.Create(userPO)
	t.ID = userID
	return err
}

// 解析登录信息
func (t *User) Parse(ctx context.Context) (info User, err error) {
	token := ctx.Value(constant.TOKEN)
	cache := adapter.NewCache()
	key := fmt.Sprintf("%s:%s", constant.TOKEN, token)
	sUserID, err := cache.Get(key)
	if err != nil || sUserID == "" {
		err = errors.New("未取得登录信息")
		return
	}
	key = fmt.Sprintf("%v", constant.USER)
	sUser := cache.HGet(key, sUserID)
	if sUser == "" {
		err = errors.New("登录信息已过期")
		return
	}

	err = json.Unmarshal([]byte(sUser), &info)
	return
}
