package goodsService

import "golang_testovoe/cmd/http_server/app/controller/goodsController/goodsService/dto"

type GoodsService struct {
}

func NewGoodsService() *GoodsService {
	return &GoodsService{}
}

func (g *GoodsService) List() {
}

func (g *GoodsService) Create(request dto.GoodsCreateRequest) (dto.GoodsResponse, error) {
	var res dto.GoodsResponse
	return res, nil
}

func (g *GoodsService) Update(request dto.GoodsUpdateRequest) (dto.GoodsResponse, error) {
	var res dto.GoodsResponse
	return res, nil
}

func (g *GoodsService) Delete(id int) error {
	return nil
}

func (g *GoodsService) GetMaxPriority(id, productId string) (int, error) {
	return 1, nil
}
