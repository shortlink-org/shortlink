package main

import (
	"fmt"

	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/infrastructure/persistence"
	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/interfaces/controller"
	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/usecase"
)

func main() {
	goodsRepo := persistence.NewGoodsJSONRepository("tests/fixtures/phone.json")
	goodsUseCase := usecase.NewGoodsUseCase(goodsRepo)
	goodsController := controller.NewGoodsController(goodsUseCase)

	err := goodsController.GenerateFeeds("policy", "out")
	if err != nil {
		fmt.Println("Error generating feeds:", err)
	}
}
