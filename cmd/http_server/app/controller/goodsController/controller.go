package goodsController

import (
	"encoding/json"
	"golang_testovoe/cmd/http_server/app/controller/goodsController/goodsService"
	"golang_testovoe/cmd/http_server/app/controller/goodsController/goodsService/dto"
	"net/http"
)

type GoodsController struct {
}

func NewGoodsController() *GoodsController {
	return &GoodsController{}
}

func (GoodsController) List(w http.ResponseWriter, r *http.Request) {

}

func (GoodsController) Create(w http.ResponseWriter, r *http.Request) {
	gService := goodsService.NewGoodsService()
	var req dto.GoodsCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusUnprocessableEntity)
		return
	}
	newGoods, err := gService.Create(req)
	if err != nil {
		http.Error(w, "Error create", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(newGoods); err != nil {
		http.Error(w, "Encode error", http.StatusInternalServerError)
		return
	}

}

func (GoodsController) Update(w http.ResponseWriter, r *http.Request) {
	gService := goodsService.NewGoodsService()
	var req dto.GoodsUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusUnprocessableEntity)
		return
	}
	updatedGoods, err := gService.Update(req)
	if err != nil {
		http.Error(w, "Error update", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(updatedGoods); err != nil {
		http.Error(w, "Encode error", http.StatusInternalServerError)
		return
	}
}

func (GoodsController) Delete(w http.ResponseWriter, r *http.Request) {

}
func (GoodsController) RePrioritize(w http.ResponseWriter, r *http.Request) {

}
