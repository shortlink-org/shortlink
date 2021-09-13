//go:generate gotext -srclang=en-GB update -out=catalog.go -lang=en-GB,de-DE,fr-CH github.com/batazor/shortlink/internal/services/api/application/http-chi

package i18n

import (
	"context"

	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func New(ctx context.Context) *message.Printer {
	// Declare variable to hold the target language tag.
	var lang language.Tag

	// Detect language
	viper.AutomaticEnv()
	viper.SetDefault("APP_LANGUAGE", "en-gb") // Select: postgres, mongo, mysql, redis, dgraph, sqlite, leveldb, badger, ram

	// Use language.MustParse() to assign the appropriate language tag
	// for the locale.
	switch viper.GetString("APP_LANGUAGE") {
	case "en-gb":
		lang = language.MustParse("en-GB")
	case "de-de":
		lang = language.MustParse("de-DE")
	case "fr-ch":
		lang = language.MustParse("fr-CH")
	default:
		lang = language.MustParse("de-DE")
	}

	// Initialize a message.Printer which uses the target language.
	p := message.NewPrinter(lang)

	return p
}
