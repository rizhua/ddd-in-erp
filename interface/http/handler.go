package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"rizhua.com/application"
	"rizhua.com/infrastructure/adapter"
	"rizhua.com/infrastructure/constant"
	"rizhua.com/infrastructure/etc"
)

// 处理成功响应
func ok(ctx *gin.Context, data interface{}) {
	if data != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"data": data,
			"desc": "Success",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"desc": "Success",
		})
	}
}

// 处理失败响应
func fail(ctx *gin.Context, code int32, desc string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"desc": desc,
	})
}

func NewHandler(
	brandApp application.BrandService,
	bundleApp application.BundleService,
	categoryApp application.CategoryService,
	nodeApp application.NodeService,
	orderApp application.OrderService,
	productApp application.ProductService,
	roleApp application.RoleService,
	structureApp application.StructureService,
	systemApp application.SystemService,
	userApp application.UserService,
) Handler {
	return Handler{
		brandApp:     brandApp,
		bundleApp:    bundleApp,
		categoryApp:  categoryApp,
		nodeApp:      nodeApp,
		orderApp:     orderApp,
		productApp:   productApp,
		roleApp:      roleApp,
		structureApp: structureApp,
		systemApp:    systemApp,
		userApp:      userApp,
	}
}

type Handler struct {
	Context      context.Context
	brandApp     application.BrandService
	bundleApp    application.BundleService
	categoryApp  application.CategoryService
	nodeApp      application.NodeService
	orderApp     application.OrderService
	structureApp application.StructureService
	productApp   application.ProductService
	roleApp      application.RoleService
	systemApp    application.SystemService
	userApp      application.UserService
}

func (t *Handler) Rest(c *gin.Context) {
	var (
		data any
		err  error
	)
	params := make(map[string]any)
	c.ShouldBindJSON(&params)
	args, err := json.Marshal(params)
	if err != nil {
		fail(c, 4010, err.Error())
		return
	}
	token := c.GetHeader("Access-Token")
	if token != "" {
		t.WithToken(token)
	}

	arr := strings.Split(c.Request.URL.Path, "/")
	if len(arr) < 2 {
		fail(c, 3050, "404")
		return
	}

	switch arr[1] {
	case "brand":
		data, err = t.Brand(c.Request.URL.Path, args)
	case "bundle":
		data, err = t.Bundle(c.Request.URL.Path, args)
	case "category":
		data, err = t.Category(c.Request.URL.Path, args)
	case "node":
		data, err = t.Node(c.Request.URL.Path, args)
	case "structure":
		data, err = t.Structure(c.Request.URL.Path, args)
	case "product":
		data, err = t.Product(c.Request.URL.Path, args)
	case "role":
		data, err = t.Role(c.Request.URL.Path, args)
	case "user":
		data, err = t.User(c.Request.URL.Path, args)
	case "order":
		data, err = t.Order(c.Request.URL.Path, args)
	case "system":
		data, err = t.System(c.Request.URL.Path, args)
	default:
		err = errors.New("404")
	}

	if err != nil {
		fail(c, 3050, err.Error())
		return
	}
	ok(c, data)
}

// 品牌
func (t *Handler) Brand(path string, args []byte) (data any, err error) {
	switch path {
	case "/brand/create":
		err = t.brandApp.Create(args)
	case "/brand/delete":
		err = t.brandApp.Delete(args)
	case "/brand/update":
		err = t.brandApp.Update(args)
	case "/brand/find":
		data, err = t.brandApp.Find(args)
	default:
		err = errors.New("404")
	}
	return
}

// 套餐
func (t *Handler) Bundle(path string, args []byte) (data any, err error) {
	switch path {
	case "/bundle/create":
		err = t.bundleApp.Create(args)
	case "/bundle/delete":
		err = t.bundleApp.Delete(args)
	case "/bundle/update":
		err = t.bundleApp.Update(args)
	case "/bundle/find":
		data, err = t.bundleApp.Find(args)
	case "/bundle/bindNodeId":
		err = t.bundleApp.BindNodeID(args)
	case "/bundle/findNodeId":
		data, err = t.bundleApp.FindNodeID(args)
	case "/bundle/license/find":
		data, err = t.bundleApp.FindLicense(args)
	default:
		err = errors.New("404")
	}
	return
}

// 商品类目
func (t *Handler) Category(path string, args []byte) (data any, err error) {
	switch path {
	case "/category/create":
		err = t.categoryApp.Create(args)
	case "/category/delete":
		err = t.categoryApp.Delete(args)
	case "/category/update":
		err = t.categoryApp.Update(args)
	case "/category/find":
		data, err = t.categoryApp.Find(args)
	case "/category/attribute/create":
		err = t.categoryApp.CreateAttribute(args)
	case "/category/attribute/delete":
		err = t.categoryApp.DeleteAttribute(args)
	case "/category/attribute/update":
		err = t.categoryApp.UpdateAttribute(args)
	case "/category/attribute/find":
		data, err = t.categoryApp.FindAttribute(args)
	default:
		err = errors.New("404")
	}
	return
}

// 节点
func (t *Handler) Node(path string, args []byte) (data any, err error) {
	switch path {
	case "/node/create":
		err = t.nodeApp.Create(args)
	case "/node/delete":
		err = t.nodeApp.Delete(args)
	case "/node/update":
		err = t.nodeApp.Update(args)
	case "/node/find":
		data, err = t.nodeApp.Find(args)
	case "/node/setSort":
		err = t.nodeApp.SetSort(args)
	case "/node/setStatus":
		err = t.nodeApp.SetStatus(args)
	case "/node/permission":
		data, err = t.nodeApp.Permission(args)
	default:
		err = errors.New("404")
	}
	return
}

// 订单
func (t *Handler) Order(path string, args []byte) (data any, err error) {
	switch path {
	case "/order/unified":
		data, err = t.orderApp.Unified(args)
	default:
		err = errors.New("404")
	}
	return
}

// 组织架构
func (t *Handler) Structure(path string, args []byte) (data any, err error) {
	switch path {
	case "/structure/org/find":
		data, err = t.structureApp.Find(args)
	case "/structure/org/switch":
		err = t.structureApp.Switch(args)
	case "/structure/dept/create":
		err = t.structureApp.CreateDept(args)
	case "/structure/dept/delete":
		err = t.structureApp.DeleteDept(args)
	case "/structure/dept/update":
		err = t.structureApp.UpdateDept(args)
	case "/structure/node/find":
		data, err = t.structureApp.FindNode()
	case "/structure/dept/find":
		data, err = t.structureApp.FindDept(args)
	case "/structure/emp/create":
		err = t.structureApp.CreateEmp(args)
	case "/structure/emp/update":
		err = t.structureApp.UpdateEmp(args)
	case "/structure/emp/find":
		data, err = t.structureApp.FindEmp(args)
	default:
		err = errors.New("404")
	}
	return
}

// 商品
func (t *Handler) Product(path string, args []byte) (data any, err error) {
	switch path {
	case "/product/create":
		err = t.productApp.Create(args)
	case "/product/delete":
		err = t.productApp.Delete(args)
	case "/product/update":
		err = t.productApp.Update(args)
	case "/product/get":
		data, err = t.productApp.Get(args)
	case "/product/find":
		data, err = t.productApp.Find(args)
	case "/product/attribute/create":
		err = t.productApp.CreateAttribute(args)
	case "/product/attribute/delete":
		err = t.productApp.DeleteAttribute(args)
	case "/product/attribute/update":
		err = t.productApp.UpdateAttribute(args)
	case "/product/attribute/find":
		data, err = t.productApp.FindAttribute(args)
	default:
		err = errors.New("404")
	}
	return
}

// 角色
func (t *Handler) Role(path string, args []byte) (data any, err error) {
	switch path {
	case "/role/create":
		err = t.roleApp.Create(args)
	case "/role/delete":
		err = t.roleApp.Delete(args)
	case "/role/update":
		err = t.roleApp.Update(args)
	case "/role/find":
		data, err = t.roleApp.Find(args)
	case "/role/bindNodeId":
		err = t.roleApp.BindNodeID(args)
	case "/role/findNodeId":
		data, err = t.roleApp.FindNodeID(args)
	case "/role/findNode":
		data, err = t.roleApp.FindNode(args)
	case "/role/addUser":
		err = t.roleApp.AddUser(args)
	case "/role/removeUser":
		err = t.roleApp.RemoveUser(args)
	case "/role/findUser":
		data, err = t.roleApp.FindUser(args)
	default:
		err = errors.New("404")
	}
	return
}

// 系统
func (t *Handler) System(path string, args []byte) (data any, err error) {
	switch path {
	case "/system/sendSms":
		err = t.systemApp.SendSms(args)
	case "/system/sendEmail":
		err = t.systemApp.SendEmail(args)
	default:
		err = errors.New("404")
	}
	return
}

// 用户
func (t *Handler) User(path string, args []byte) (data any, err error) {
	switch path {
	case "/user/signIn":
		data, err = t.userApp.SignIn(args)
	case "/user/signUp":
		data, err = t.userApp.SignUp(args)
	case "/user/active":
		err = t.userApp.Active(args)
	case "/user/forget":
		data, err = t.userApp.Forget(args)
	case "/user/find":
		data, err = t.userApp.Find(args)
	case "/user/rePassword":
		err = t.userApp.RePassword(args)
	case "/user/setPassword":
		err = t.userApp.SetPassword(args)
	case "/user/work":
		data, err = t.userApp.Work()
	case "/user/parse":
		data, err = t.userApp.Parse()
	default:
		err = errors.New("404")
	}
	return
}

// 退出登录
func (t *Handler) Logout(c *gin.Context) {
	token := c.GetHeader("Access-Token")
	cache := adapter.NewCache()
	key := fmt.Sprintf("%s:%s", constant.TOKEN, token)
	err := cache.Del(key)
	if err != nil {
		fail(c, 4030, "退出登录失败")
	}

	ok(c, nil)
}

// 上传
func (t *Handler) Upload(ctx *gin.Context) {
	var (
		dst interface{}
		err error
	)

	batch := ctx.PostForm("batch")
	if batch == "true" {
		t.multi(ctx)
	} else {
		dst, err = t.single(ctx)
	}

	if err != nil {
		fail(ctx, 4011, err.Error())
		return
	}

	ok(ctx, dst)
}

// 单文件上传
func (t *Handler) single(ctx *gin.Context) (dst string, err error) {
	f, _ := ctx.FormFile("file")

	if !t.checkSize(f) {
		fail(ctx, 4011, "The file exceeds the specified size")
		return
	}
	if !t.checkExt(f) {
		fail(ctx, 4012, "Invalid file format")
		return
	}

	cfg := etc.C
	now := time.Now()
	ymd := fmt.Sprintf("%d/%d/%d/", now.Year(), now.Month(), now.Day())
	fp := cfg.File.Path + ymd
	os.Mkdir(fp, os.ModePerm)
	dst = path.Join(fp, f.Filename)
	err = ctx.SaveUploadedFile(f, dst)
	if err != nil {
		fail(ctx, 4013, "Save File Fail")
	}
	dst = "/" + cfg.File.Prefix + "/" + ymd + f.Filename
	return
}

// 多文件上传
func (t *Handler) multi(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return
	}
	files := form.File["file"]
	for _, v := range files {

		fmt.Println(v.Filename)
	}
}

// 校验类型
func (t *Handler) checkExt(f *multipart.FileHeader) bool {
	var flag bool

	fileType := f.Header.Get("Content-Type")
	ext := [...]string{"image/jpeg", "image/png"}
	for _, v := range ext {
		if v == fileType {
			flag = true
		}
	}

	return flag
}

// 校验大小
func (t *Handler) checkSize(f *multipart.FileHeader) bool {
	cfg := etc.C
	return f.Size < cfg.File.Size
}

func (t *Handler) WithToken(token string) {
	ctx := context.WithValue(context.Background(), constant.TOKEN, token)
	t.brandApp.Context = ctx
	t.bundleApp.Context = ctx
	t.categoryApp.Context = ctx
	t.nodeApp.Context = ctx
	t.orderApp.Context = ctx
	t.structureApp.Context = ctx
	t.productApp.Context = ctx
	t.roleApp.Context = ctx
	t.systemApp.Context = ctx
	t.userApp.Context = ctx
}
