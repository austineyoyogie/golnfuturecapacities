package routes

import (
	"golnfuturecapacities/api/config/database"
	"golnfuturecapacities/api/resource"
	"golnfuturecapacities/api/service/implementation"
)

var (
	CD                 = database.CDriver()
	welcomeServiceImpl = implementation.NewWelcomeServiceImpl(CD)
	welcomeResource    = resource.WelcomeController(welcomeServiceImpl)
	/*----------------------------------------------------*/
	userServiceImpl = implementation.NewUserServiceImpl(CD)
	userResource    = resource.UserController(userServiceImpl)
	/*----------------------------------------------------*/
	roleServiceImpl = implementation.NewRoleServiceImpl(CD)
	roleResource    = resource.RoleController(roleServiceImpl)
	/*----------------------------------------------------*/
	supplyServiceImpl = implementation.NewSupplyServiceImpl(CD)
	supplyResource    = resource.SupplyController(supplyServiceImpl)
	/*----------------------------------------------------*/
	categoryServiceImpl = implementation.NewCategoryServiceImpl(CD)
	categoryResource    = resource.CategoryController(categoryServiceImpl)
	/*----------------------------------------------------*/
	productServiceImpl = implementation.NewProductServiceImpl(CD)
	productResource    = resource.ProductController(productServiceImpl)
)
