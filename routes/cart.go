package routes

import (
	"bewaysbuck/handlers"
	"bewaysbuck/pkg/middleware"
	"bewaysbuck/pkg/mysql"
	"bewaysbuck/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	CartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(CartRepository)

	r.HandleFunc("/carts", middleware.Auth(h.FindCarts)).Methods("GET")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.GetCart)).Methods("GET")
	r.HandleFunc("/cart", middleware.Auth(middleware.UploadFile(h.CreateCart))).Methods("POST")
	r.HandleFunc("/cart/{id}", middleware.Auth(middleware.UploadFile(h.UpdateCart))).Methods("PATCH")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.DeleteCart)).Methods("DELETE")
}
