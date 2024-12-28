package application

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
	"time"

	"golang.org/x/exp/rand"
	"rizhua.com/application/command"
	"rizhua.com/domain/entity"
	"rizhua.com/infrastructure/adapter"
	"rizhua.com/infrastructure/constant"
	"rizhua.com/infrastructure/etc"
)

func NewSystemService() SystemService {
	return SystemService{}
}

type SystemService struct {
	Context context.Context
}

func (c *SystemService) SendSms(args []byte) error {
	cmd := command.SendSms{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	code := rand.New(rand.NewSource(uint64(time.Now().UnixNano()))).Int31n(10000)
	cache := adapter.NewCache()
	key := fmt.Sprintf("%s:%s", constant.CAPTCHA, cmd.Mobile)
	return cache.Set(key, code, 30*time.Minute)
}

func (c *SystemService) SendEmail(args []byte) error {
	cmd := command.SendEmail{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	cfg := etc.C
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	from := mail.Address{Name: "日抓云", Address: cfg.Email.User} // 发件人
	to := mail.Address{Name: "", Address: cmd.Email}           // 收件人

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(cmd.TemplateParam["subject"])))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"

	// 获取模板
	configDO := entity.ConfigEntity{}
	config, err := configDO.GetByCode(cmd.TemplateCode)
	if err != nil {
		return err
	}
	body := ""
	if ok := config.Data["template"]; ok == nil {
		return fmt.Errorf("模板不存在")
	} else {
		body = config.Data["template"].(string)
	}
	for k, v := range cmd.TemplateParam {
		body = strings.Replace(body, fmt.Sprintf("{{%s}}", k), v, -1)
	}

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))

	hp := strings.Split(cfg.Email.Host, ":")
	auth := smtp.PlainAuth("", cfg.Email.User, cfg.Email.Password, hp[0])
	return smtp.SendMail(cfg.Email.Host, auth, cfg.Email.User, []string{to.Address}, []byte(message))
}
