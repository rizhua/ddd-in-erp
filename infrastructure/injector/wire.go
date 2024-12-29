//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"github.com/google/wire"
	"rizhua.com/application"
	"rizhua.com/domain"
	"rizhua.com/infrastructure/persistence"
	"rizhua.com/interface/http"
	"xorm.io/xorm"
)

func BuildInjector(db *xorm.Engine) (*Injector, func(), error) {
	wire.Build(
		InitGinEngine,
		// InitCasbin,

		// infrastructure
		persistence.NewAttribute,
		persistence.NewBrand,
		persistence.NewBundle,
		persistence.NewCategory,
		persistence.NewDept,
		persistence.NewEmp,
		persistence.NewNode,
		persistence.NewOrder,
		persistence.NewOrg,
		persistence.NewRole,
		persistence.NewSpu,
		persistence.NewUser,

		// domain
		domain.NewBrandService,
		domain.NewBundleService,
		domain.NewCategoryService,
		domain.NewNodeService,
		domain.NewOrderService,
		domain.NewProductService,
		domain.NewRoleService,
		domain.NewStructureService,
		domain.NewUserService,

		// application
		application.NewBrandService,
		application.NewBundleService,
		application.NewCategoryService,
		application.NewNodeService,
		application.NewOrderService,
		application.NewProductService,
		application.NewRoleService,
		application.NewStructureService,
		application.NewSystemService,
		application.NewUserService,

		// interface
		http.NewHandler,

		// router
		http.NewRouter,

		NewInjector,
	)
	return &Injector{}, nil, nil
}
