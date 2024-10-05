package persistence

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/domain/entity"
	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/domain/repository"
)

type GoodsJSONRepository struct {
	FilePath string
}

func NewGoodsJSONRepository(filePath string) repository.GoodsRepository {
	return &GoodsJSONRepository{FilePath: filePath}
}

func (repo *GoodsJSONRepository) GetAllGoods() ([]entity.Goods, error) {
	var goods []entity.Goods
	data, err := os.ReadFile(repo.FilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading goods data: %w", err)
	}
	err = json.Unmarshal(data, &goods)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling goods data: %w", err)
	}
	return goods, nil
}
