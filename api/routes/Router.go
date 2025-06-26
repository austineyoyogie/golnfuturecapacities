package routes

import (
	"net/http"
)

func RouteHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", mux))
	mux.HandleFunc("GET /welcome", welcomeResource.WelcomeGetHandler)
	/*-------------------------------------------------------------*/
	mux.HandleFunc("POST /login", userResource.LoginUserHandler)
	mux.HandleFunc("POST /register", userResource.RegisterUserHandler)
	mux.HandleFunc("GET /profile/{id}", userResource.FindUserIdHandler)
	mux.HandleFunc("GET /profile", userResource.FindAllUsersHandler)
	mux.HandleFunc("PUT /profile/{id}", userResource.UpdateUserIdHandler)
	mux.HandleFunc("DELETE /user/{id}", userResource.DeleteUserIdHandler)
	/*-------------------------------------------------------------*/
	mux.HandleFunc("POST /role", roleResource.CreateRoleHandler)
	mux.HandleFunc("GET /role/{id}", roleResource.FindRoleIdHandler)
	mux.HandleFunc("GET /role", roleResource.FindAllRoleHandler)
	mux.HandleFunc("PUT /role/{id}", roleResource.UpdateRoleIdHandler)
	mux.HandleFunc("DELETE /role/{id}", roleResource.DeleteRoleIdHandler)
	/*-------------------------------------------------------------*/
	mux.HandleFunc("POST /supply", supplyResource.CreateSupplyHandler)
	mux.HandleFunc("GET /supply/{id}", supplyResource.FindSupplyIdHandler)
	mux.HandleFunc("GET /supplier", supplyResource.FindAllSupplyHandler)
	mux.HandleFunc("PUT /supply/{id}", supplyResource.UpdateSupplyIdHandler)
	mux.HandleFunc("DELETE /supply/{id}", supplyResource.DeleteSupplyIdHandler)
	/*-------------------------------------------------------------*/
	mux.HandleFunc("POST /category", categoryResource.CreateCategoryHandler)
	mux.HandleFunc("GET /category/{id}", categoryResource.FindCategoryIdHandler)
	mux.HandleFunc("GET /categories", categoryResource.FindAllCategoryHandler)
	mux.HandleFunc("PUT /category/{id}", categoryResource.UpdateCategoryIdHandler)
	mux.HandleFunc("DELETE /category/{id}", categoryResource.DeleteCategoryIdHandler)
	/*-------------------------------------------------------------*/
	mux.HandleFunc("POST /product", productResource.CreateProductHandler)
	mux.HandleFunc("GET /product/{id}", productResource.FindProductIdHandler)
	mux.HandleFunc("GET /products", productResource.FindAllProductHandler)
	mux.HandleFunc("PUT /product/{id}", productResource.UpdateProductIdHandler)
	mux.HandleFunc("DELETE /product/{id}", productResource.DeleteProductIdHandler)
	return mux
}
