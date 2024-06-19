package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/roh4nyh/ecom/service/cart"
	"github.com/roh4nyh/ecom/service/order"
	"github.com/roh4nyh/ecom/service/product"
	"github.com/roh4nyh/ecom/service/user"
	"github.com/roh4nyh/ecom/utils"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		db:   db,
		addr: addr,
	}
}

func (s *ApiServer) Run() error {

	router := http.NewServeMux()
	server := http.Server{
		Addr:         s.addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	subRouter := http.NewServeMux()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", subRouter))

	// Register user handler...
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	// Register product handler...
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subRouter)

	// Register cart handler...
	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subRouter)

	// health check
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"status": "server is up and running..."}
		utils.WriteJSON(w, http.StatusOK, response)
	})

	return server.ListenAndServe()
}
