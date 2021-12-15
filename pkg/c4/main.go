package main

import (
	"fmt"
	"os"

	"github.com/krzysztofreczek/go-structurizr/pkg/scraper"
	"github.com/krzysztofreczek/go-structurizr/pkg/view"

	"github.com/batazor/shortlink/internal/di"
)

const (
	scraperConfig = "./pkg/c4/scraper.yml"
	viewConfig    = "./pkg/c4/view.yml"
	outputFile    = "./pkg/c4/out/view-%s.plantuml"
)

func main() {
	// Init services
	serviceAPIService, _, err := di.InitializeAPIService()
	if err != nil {
		panic(err)
	}
	scrape(serviceAPIService, "APIService")

	serviceBillingService, _, err := di.InitializeBillingService()
	if err != nil {
		panic(err)
	}
	scrape(serviceBillingService, "BillingService")

	serviceLinkService, _, err := di.InitializeLinkService()
	if err != nil {
		panic(err)
	}
	scrape(serviceLinkService, "LinkService")

	serviceLoggerService, _, err := di.InitializeLoggerService()
	if err != nil {
		panic(err)
	}
	scrape(serviceLoggerService, "LoggerService")

	serviceMetadataService, _, err := di.InitializeMetadataService()
	if err != nil {
		panic(err)
	}
	scrape(serviceMetadataService, "MetadataService")

	serviceNotifyService, _, err := di.InitializeNotifyService()
	if err != nil {
		panic(err)
	}
	scrape(serviceNotifyService, "NotifyService")
}

func scrape(app interface{}, name string) {
	s, err := scraper.NewScraperFromConfigFile(scraperConfig)
	if err != nil {
		panic(err)
	}

	structure := s.Scrape(app)

	outFileName := fmt.Sprintf(outputFile, name)
	outFile, err := os.Create(outFileName)
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
