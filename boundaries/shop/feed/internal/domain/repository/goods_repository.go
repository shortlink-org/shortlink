package repository

import (
	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/domain/entity"
)

type GoodsRepository interface {
	GetAllGoods() ([]entity.Goods, error)
}
