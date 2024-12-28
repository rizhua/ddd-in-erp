package http

import (
	"github.com/gin-gonic/gin"
	"rizhua.com/interface/middleware"
)

func NewRouter(
	Handler Handler,
) *Router {
	return &Router{
		handler: Handler,
	}
}

type Router struct {
	handler Handler
}

func (r *Router) Register(app *gin.Engine) error {
	brand := app.Group("/brand")
	{
		brand.POST("/find", r.handler.Rest)
		brand.Use(middleware.Auth())

		brand.POST("/create", r.handler.Rest)
		brand.POST("/delete", r.handler.Rest)
		brand.POST("/update", r.handler.Rest)
	}

	bundle := app.Group("/bundle")
	{
		bundle.POST("/find", r.handler.Rest)

		bundle.Use(middleware.Auth())

		bundle.POST("/create", r.handler.Rest)
		bundle.POST("/delete", r.handler.Rest)
		bundle.POST("/update", r.handler.Rest)
		bundle.POST("/bindNodeId", r.handler.Rest)
		bundle.POST("/findNodeId", r.handler.Rest)
		bundle.POST("/license/find", r.handler.Rest)
	}

	file := app.Group("/file")
	{
		file.POST("/upload", r.handler.Upload)
	}

	node := app.Group("/node")
	{
		node.Use(middleware.Auth())

		node.POST("/create", r.handler.Rest)
		node.POST("/delete", r.handler.Rest)
		node.POST("/update", r.handler.Rest)
		node.POST("/find", r.handler.Rest)
		node.POST("/setSort", r.handler.Rest)
		node.POST("/setStatus", r.handler.Rest)
		node.POST("/permission", r.handler.Rest)
	}

	order := app.Group("/order")
	{
		order.POST("/unified", r.handler.Rest)
	}

	structure := app.Group("/structure")
	{
		structure.Use(middleware.Auth())

		structure.POST("/org/find", r.handler.Rest)
		structure.POST("/org/switch", r.handler.Rest)
		structure.POST("/node/find", r.handler.Rest)

		structure.POST("/dept/create", r.handler.Rest)
		structure.POST("/dept/delete", r.handler.Rest)
		structure.POST("/dept/update", r.handler.Rest)
		structure.POST("/dept/find", r.handler.Rest)

		structure.POST("/emp/create", r.handler.Rest)
		structure.POST("/emp/update", r.handler.Rest)
		structure.POST("/emp/find", r.handler.Rest)
	}

	product := app.Group("/product")
	{
		product.POST("/info", r.handler.Rest)
		product.POST("/find", r.handler.Rest)
		product.POST("/brand/find", r.handler.Rest)

		product.Use(middleware.Auth())

		product.POST("/create", r.handler.Rest)
		product.POST("/delete", r.handler.Rest)
		product.POST("/update", r.handler.Rest)
		product.POST("/attribute/create", r.handler.Rest)
		product.POST("/attribute/delete", r.handler.Rest)
		product.POST("/attribute/update", r.handler.Rest)
		product.POST("/attribute/find", r.handler.Rest)
		product.POST("/category/create", r.handler.Rest)
		product.POST("/category/delete", r.handler.Rest)
		product.POST("/category/update", r.handler.Rest)
		product.POST("/category/find", r.handler.Rest)
	}

	role := app.Group("/role")
	{
		role.Use(middleware.Auth())

		role.POST("/create", r.handler.Rest)
		role.POST("/delete", r.handler.Rest)
		role.POST("/update", r.handler.Rest)
		role.POST("/find", r.handler.Rest)
		role.POST("/bindNodeId", r.handler.Rest)
		role.POST("/findNodeId", r.handler.Rest)
		role.POST("/findNode", r.handler.Rest)
		role.POST("/addUser", r.handler.Rest)
		role.POST("/removeUser", r.handler.Rest)
		role.POST("/findUser", r.handler.Rest)
	}

	sys := app.Group("/system")
	{
		sys.POST("/sms/send", r.handler.Rest)
		sys.POST("/email/send", r.handler.Rest)

		sys.Use(middleware.Auth())
	}

	user := app.Group("/user")
	{
		user.POST("/signIn", r.handler.Rest)
		user.POST("/signUp", r.handler.Rest)
		user.POST("/active", r.handler.Rest)
		user.POST("/forget", r.handler.Rest)
		user.POST("/password/reset", r.handler.Rest)

		user.Use(middleware.Auth())

		user.POST("/work", r.handler.Rest)
		user.POST("/parse", r.handler.Rest)
		user.POST("/find", r.handler.Rest)
		user.POST("/setPassword", r.handler.Rest)
		user.POST("/logout", r.handler.Logout)
	}

	return nil
}
