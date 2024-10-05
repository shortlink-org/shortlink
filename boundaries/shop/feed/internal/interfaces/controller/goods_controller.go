package controller

import (
	"fmt"

	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/usecase"
	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/utils"
)

type GoodsController struct {
	GoodsUseCase usecase.GoodsUseCase
}

func NewGoodsController(uc usecase.GoodsUseCase) *GoodsController {
	return &GoodsController{
		GoodsUseCase: uc,
	}
}

func (gc *GoodsController) GenerateFeeds(policyDir, outputDir string) error {
	// Load policies
	policyFiles, err := utils.GetPolicyFiles(policyDir)
	if err != nil {
		return err
	}

	for _, policyFile := range policyFiles {
		policy, err := utils.LoadPolicy(policyFile)
		if err != nil {
			return err
		}

		filteredGoods, err := gc.GoodsUseCase.GetFilteredGoods(policy)
		if err != nil {
			return err
		}

		if len(filteredGoods) > 0 {
			outputFileName := fmt.Sprintf("%s/feed_%s.xml", outputDir, policy.Name)
			err = utils.GenerateXML(filteredGoods, outputFileName)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("No goods passed the policy %s, skipping XML generation.\n", policy.Name)
		}
	}
	return nil
}
