package usecase

import (
	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/domain/entity"
	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/domain/repository"
)

type GoodsUseCase interface {
	GetFilteredGoods(policy Policy) ([]entity.Goods, error)
}

type goodsUseCase struct {
	goodsRepo repository.GoodsRepository
}

func NewGoodsUseCase(repo repository.GoodsRepository) GoodsUseCase {
	return &goodsUseCase{
		goodsRepo: repo,
	}
}

func (uc *goodsUseCase) GetFilteredGoods(policy Policy) ([]entity.Goods, error) {
	goods, err := uc.goodsRepo.GetAllGoods()
	if err != nil {
		return nil, err
	}

	var filteredGoods []entity.Goods
	for _, item := range goods {
		match, err := EvaluatePolicy(policy, item)
		if err != nil {
			return nil, err
		}
		if match {
			filteredGoods = append(filteredGoods, item)
		}
	}
	return filteredGoods, nil
}
