package utils

import (
	"encoding/xml"
	"fmt"
	"os"
	"sort"

	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/domain/entity"
)

type Feature struct {
	XMLName xml.Name `xml:"feature"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",chardata"`
}

type XMLGoods struct {
	XMLName  xml.Name  `xml:"goods"`
	Brand    string    `xml:"brand"`
	Model    string    `xml:"model"`
	Price    string    `xml:"price"`
	Stock    int       `xml:"stock"`
	Category string    `xml:"category"`
	Tags     []string  `xml:"tags>tag"`
	Features []Feature `xml:"features>feature"`
}

func GenerateXML(goods []entity.Goods, filePath string) error {
	var xmlGoodsList []XMLGoods
	for _, g := range goods {
		// Sort features
		featureNames := make([]string, 0, len(g.Features))
		for k := range g.Features {
			featureNames = append(featureNames, k)
		}
		sort.Strings(featureNames)

		features := make([]Feature, 0, len(featureNames))
		for _, k := range featureNames {
			v := g.Features[k]
			valueStr := fmt.Sprintf("%v", v)
			features = append(features, Feature{
				Name:  k,
				Value: valueStr,
			})
		}

		// Sort tags
		sort.Strings(g.Tags)

		xmlGoods := XMLGoods{
			Brand:    g.Brand,
			Model:    g.Model,
			Price:    g.Price.StringFixed(2),
			Stock:    g.Stock,
			Category: g.Category,
			Tags:     g.Tags,
			Features: features,
		}
		xmlGoodsList = append(xmlGoodsList, xmlGoods)
	}

	feed := struct {
		XMLName xml.Name   `xml:"feed"`
		Goods   []XMLGoods `xml:"goods"`
	}{
		Goods: xmlGoodsList,
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	err = encoder.Encode(feed)
	if err != nil {
		return fmt.Errorf("error encoding XML: %w", err)
	}
	return nil
}
