package main

import (
	"fmt"
	"os"

	api_di "github.com/batazor/shortlink/internal/services/api/di"
	billing_di "github.com/batazor/shortlink/internal/services/billing/di"
	link_di "github.com/batazor/shortlink/internal/services/link/di"
	logger_di "github.com/batazor/shortlink/internal/services/logger/di"
	metadata_di "github.com/batazor/shortlink/internal/services/metadata/di"
	"github.com/krzysztofreczek/go-structurizr/pkg/scraper"
	"github.com/krzysztofreczek/go-structurizr/pkg/view"
)

const (
	scraperConfig = "./pkg/c4/scraper.yml"
	viewConfig    = "./pkg/c4/view.yml"
	outputFile    = "./docs/c4/view-%s.plantuml"
)

func main() {
	// Init services
	serviceAPIService, _, err := api_di.InitializeAPIService()
	if err != nil {
		panic(err)
	}
	scrape(serviceAPIService, "APIService")

	serviceBillingService, _, err := billing_di.InitializeBillingService()
	if err != nil {
		panic(err)
	}
	scrape(serviceBillingService, "BillingService")

	serviceLinkService, _, err := link_di.InitializeLinkService()
	if err != nil {
		panic(err)
	}
	scrape(serviceLinkService, "LinkService")

	serviceLoggerService, _, err := logger_di.InitializeLoggerService()
	if err != nil {
		panic(err)
	}
	scrape(serviceLoggerService, "LoggerService")

	serviceMetadataService, _, err := metadata_di.InitializeMetaDataService()
	if err != nil {
		panic(err)
	}
	scrape(serviceMetadataService, "MetadataService")
}

func scrape(app interface{}, name string) {
	s, err := scraper.NewScraperFromConfigFile(scraperConfig)
	if err != nil {
		panic(err)
	}

	structure := s.Scrape(app)

	outFileName := fmt.Sprintf(outputFile, name)
	outFile, err := os.Create(outFileName) // #nosec
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = outFile.Close()
	}()

	v, err := view.NewViewFromConfigFile(viewConfig)
	if err != nil {
		panic(err)
	}

	err = v.RenderStructureTo(structure, outFile)
	if err != nil {
		panic(err)
	}
}
