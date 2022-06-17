package api

import (
	"fmt"
	"net/http"

	"github.com/rg-km/final-project-engineering-51/backend/repository"
)

type API struct {
	usersRepo repository.UserRepository
	mux             *http.ServeMux
}

func NewAPI(usersRepo repository.UserRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo,mux,
	}

	mux.Handle("/api/user/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/user/logout", api.POST(http.HandlerFunc(api.logout)))
	mux.Handle("/api/user/register", api.POST(http.HandlerFunc(api.register)))

	// API with AuthMiddleware:
	// mux.Handle("/api/products", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.productList))))
	// mux.Handle("/api/cart/add", api.POST(api.AuthMiddleWare(http.HandlerFunc(api.addToCart))))
	// mux.Handle("/api/cart/clear", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.clearCart))))
	// mux.Handle("/api/carts", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.cartList))))
	// mux.Handle("/api/pay", api.POST(api.AuthMiddleWare(http.HandlerFunc(api.pay))))

	// // API with AuthMiddleware and AdminMiddleware
	mux.Handle("/api/admin/users", api.GET(api.AuthMiddleWare(api.AdminMiddleware(http.HandlerFunc(api.userList)))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}