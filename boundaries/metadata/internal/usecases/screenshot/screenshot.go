/*
Metadata Service. Application layer
*/

package screenshot

import (
	"context"
	"net/url"

	"github.com/chromedp/chromedp"

	domainerrors "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/errors"
	s3Repository "github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/repository/media"
)

const (
	opScreenshotStoreGet = "metadata.screenshot.store.get"
	opScreenshotStorePut = "metadata.screenshot.store.put"
)

func New(ctx context.Context, media *s3Repository.Service) (*UC, error) {
	return &UC{
		media: media,
	}, nil
}

func (s *UC) Get(ctx context.Context, linkURL string) (*url.URL, error) {
	result, err := s.media.Get(ctx, linkURL)
	if err != nil {
		return nil, domainerrors.Normalize(opScreenshotStoreGet, err)
	}

	return result, nil
}

func (s *UC) Set(ctx context.Context, linkURL string) error {
	chromedp.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4585.0 Safari/537.36")

	newCtx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	// capture screenshot of an element
	var screenshot []byte

	if err := chromedp.Run(newCtx, elementScreenshot(linkURL, &screenshot)); err != nil {
		return domainerrors.NewScreenshotUnavailableError(linkURL, err)
	}

	if err := s.media.Put(ctx, linkURL, screenshot); err != nil {
		return domainerrors.Normalize(opScreenshotStorePut, err)
	}

	return nil
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EmulateViewport(defaultWidth, defaultHeight),
		chromedp.Navigate(urlstr),
		chromedp.CaptureScreenshot(res),
	}
}
