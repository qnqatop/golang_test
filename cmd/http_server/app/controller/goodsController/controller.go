package goodsController

import (
	"encoding/json"
	"golang_testovoe/cmd/http_server/app/controller/goodsController/goodsService"
	"golang_testovoe/cmd/http_server/app/controller/goodsController/goodsService/dto"
	"net/http"
)

type GoodsController struct {
	gService goodsService.GoodsService
}

func NewGoodsController() *GoodsController {
	return &GoodsController{}
}

func (g *GoodsController) List(w http.ResponseWriter, r *http.Request) {

}
func (g *GoodsController) Create(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query() // Получение Query Parameters
	projectId := queryParams.Get("projectId")
	var req dto.GoodsCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusUnprocessableEntity)
		return
	}
	newGoods, err := g.gService.Create(req, projectId)

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
func (g *GoodsController) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.GoodsUpdateRequest

	queryParams := r.URL.Query() // Получение Query Parameters
	projectId := queryParams.Get("projectId")
	goodsId := queryParams.Get("id")
	if goodsId == "" || projectId == "" {
		http.Error(w, "Bad request", http.StatusUnprocessableEntity)
		return
	}

	val, _ := g.gService.GetValueRedis(goodsId)
	if val == "" {
		http.Error(w, "errors.good.notFound", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusUnprocessableEntity)
		return
	}
	if req.Name == "" {
		http.Error(w, "Property 'name' can't be empty", http.StatusUnprocessableEntity)
		return
	}
	updatedGoods, err := g.gService.Update(req, projectId, goodsId)
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
func (g *GoodsController) Delete(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query() // Получение Query Parameters
	projectId := queryParams.Get("projectId")
	goodsId := queryParams.Get("id")
	if goodsId == "" || projectId == "" {
		http.Error(w, "Bad request", http.StatusUnprocessableEntity)
		return
	}
	val, _ := g.gService.GetValueRedis(goodsId)
	if val == "" {
		http.Error(w, "errors.good.notFound", http.StatusNotFound)
		return
	}
	res, err := g.gService.Delete(goodsId, projectId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Encode error", http.StatusInternalServerError)
		return
	}

}
func (GoodsController) RePrioritize(w http.ResponseWriter, r *http.Request) {

}
