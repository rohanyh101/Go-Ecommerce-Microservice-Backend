package product

import (
	"fmt"
	"net/http"

	"github.com/roh4nyh/ecom/service/auth"
	"github.com/roh4nyh/ecom/types"
	"github.com/roh4nyh/ecom/utils"
)

var products = []types.CreateProductPayload{
	{Name: "Parle-G", Description: "Classic glucose biscuits", Image: "parle-g.png", Price: 5.00, Quantity: 1},
	{Name: "Maggi Noodles", Description: "Instant noodles in 2 minutes", Image: "maggi.png", Price: 12.00, Quantity: 2},
	{Name: "Lays Potato Chips", Description: "Salted potato chips", Image: "lays.png", Price: 10.00, Quantity: 45},
	{Name: "Cadbury Dairy Milk", Description: "Milk chocolate bar", Image: "cadbury.png", Price: 50.00, Quantity: 2},
	{Name: "Oreo Cookies", Description: "Cream-filled chocolate cookies", Image: "oreo.png", Price: 20.00, Quantity: 76},
	{Name: "Bingo Mad Angles", Description: "Spicy potato chips", Image: "bingo.png", Price: 15.00, Quantity: 87},
	{Name: "Kurkure", Description: "Spicy corn puffs", Image: "kurkure.png", Price: 10.00, Quantity: 31},
	{Name: "Amul Butter", Description: "Salted butter", Image: "amul-butter.png", Price: 30.00, Quantity: 19},
	{Name: "Britannia NutriChoice", Description: "Digestive biscuits", Image: "nutrichoice.png", Price: 25.00, Quantity: 64},
	{Name: "Haldiram's Sev Usal", Description: "Spicy chickpea snack", Image: "sev-usal.png", Price: 20.00, Quantity: 43},
	{Name: "Tata Salt", Description: "Iodized salt", Image: "tata-salt.png", Price: 15.00, Quantity: 22},
	{Name: "Amul Cheese Slices", Description: "Processed cheese slices", Image: "amul-cheese.png", Price: 40.00, Quantity: 1},
	{Name: "Sunfeast Dark Fantasy", Description: "Chocolate-filled cookies", Image: "dark-fantasy.png", Price: 30.00, Quantity: 1},
	{Name: "Lay's Maxx", Description: "Thick-cut potato chips", Image: "lays-maxx.png", Price: 15.00, Quantity: 45},
	{Name: "Bournvita", Description: "Chocolate health drink", Image: "bournvita.png", Price: 200.00, Quantity: 21},
	{Name: "Haldiram's Aloo Bhujia", Description: "Spicy potato snack", Image: "aloo-bhujia.png", Price: 25.00, Quantity: 13},
	{Name: "Amul Ghee", Description: "Clarified butter", Image: "amul-ghee.png", Price: 350.00, Quantity: 18},
	{Name: "Britannia Good Day", Description: "Butter cookies", Image: "good-day.png", Price: 20.00, Quantity: 11},
	{Name: "Haldiram's Namkeen", Description: "Assorted savory snacks", Image: "namkeen.png", Price: 30.00, Quantity: 15},
	{Name: "Tata Sampann Urad Dal", Description: "Split black lentils", Image: "urad-dal.png", Price: 80.00, Quantity: 45},
	{Name: "Amul Kool", Description: "Flavored milk drink", Image: "amul-kool.png", Price: 25.00, Quantity: 12},
	{Name: "Sunfeast Marie Light", Description: "Light tea biscuits", Image: "marie-light.png", Price: 15.00, Quantity: 67},
	{Name: "Lay's Sizzler", Description: "Spicy potato chips", Image: "lays-sizzler.png", Price: 20.00, Quantity: 12},
	{Name: "Bournvita Pro-Health", Description: "Chocolate health drink", Image: "bournvita-pro-health.png", Price: 250.00, Quantity: 56},
	{Name: "Haldiram's Bhujia Sev", Description: "Spicy chickpea noodles", Image: "bhujia-sev.png", Price: 20.00, Quantity: 87},
	{Name: "Amul Lassi", Description: "Flavored yogurt drink", Image: "amul-lassi.png", Price: 30.00, Quantity: 12},
	{Name: "Britannia 50-50", Description: "Savory crackers", Image: "50-50.png", Price: 10.00, Quantity: 19},
	{Name: "Haldiram's Moong Dal", Description: "Spicy split green gram", Image: "moong-dal.png", Price: 25.00, Quantity: 67},
	{Name: "Tata Sampann Toor Dal", Description: "Split pigeon peas", Image: "toor-dal.png", Price: 90.00, Quantity: 23},
	{Name: "Amul Masti Dahi", Description: "Plain yogurt", Image: "amul-masti-dahi.png", Price: 20.00, Quantity: 54},
	{Name: "Sunfeast Farmlite", Description: "Digestive biscuits", Image: "farmlite.png", Price: 25.00, Quantity: 111},
	{Name: "Lay's Maxx Masala", Description: "Spicy thick-cut potato chips", Image: "lays-maxx-masala.png", Price: 20.00, Quantity: 34},
	{Name: "Bournvita Lil Champs", Description: "Chocolate health drink for kids", Image: "lil-champs.png", Price: 150.00, Quantity: 77},
	{Name: "Haldiram's Sev Usal Masala", Description: "Spicy chickpea snack mix", Image: "sev-usal-masala.png", Price: 30.00, Quantity: 32},
	{Name: "Amul Cheese Cubes", Description: "Processed cheese cubes", Image: "amul-cheese-cubes.png", Price: 45.00, Quantity: 8},
	{Name: "Britannia Milk Bikis", Description: "Milk-based cookies", Image: "milk-bikis.png", Price: 15.00, Quantity: 23},
	{Name: "Haldiram's Soya Chips", Description: "Soy-based potato chips", Image: "soya-chips.png", Price: 20.00, Quantity: 23},
	{Name: "Tata Sampann Chana Dal", Description: "Split Bengal gram", Image: "chana-dal.png", Price: 85.00, Quantity: 89},
	{Name: "Amul Butter Milk", Description: "Spiced yogurt drink", Image: "amul-butter-milk.png", Price: 25.00, Quantity: 46},
	{Name: "Sunfeast Bounce", Description: "Cream-filled cookies", Image: "bounce.png", Price: 30.00, Quantity: 56},
}

type Handler struct {
	store     types.ProductStore
	userStore types.UserStore
}

func NewHandler(s types.ProductStore, u types.UserStore) *Handler {
	return &Handler{
		store:     s,
		userStore: u,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /products", h.handleGetProducts)
	router.HandleFunc("POST /products", auth.WithJWTAuth(h.handleCreateProduct, h.userStore))

	router.HandleFunc("POST /products/bulkinsert", auth.WithJWTAuth(h.handleBunkInsert, h.userStore))
}

func (h *Handler) handleBunkInsert(w http.ResponseWriter, r *http.Request) {
	go func() {
		for _, product := range products {
			err := h.store.CreateProduct(product)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create product: %v", err))
				return
			}
		}
	}()

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "bulk data inserted..."})
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get products: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreateProduct(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create product: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "product created..."})
}
