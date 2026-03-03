/*
Metadata Service. Application layer
*/

package screenshot

import (
	"context"
	"fmt"
	"net/url"
	"time"

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
	// Add timeout for screenshot operation (30 seconds); attach cause for diagnostics
	screenshotCtx, cancel := context.WithTimeoutCause(ctx, 30*time.Second,
		fmt.Errorf("screenshot for %s: 30s timeout exceeded", linkURL))
	defer cancel()

	allocatorOpts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("user-data-dir", "/tmp/chromedp"),
		chromedp.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4585.0 Safari/537.36"),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(screenshotCtx, allocatorOpts...)
	defer cancelAlloc()

	newCtx, cancelChrome := chromedp.NewContext(allocCtx)
	defer cancelChrome()

	// capture screenshot of an element
	var screenshot []byte

	if err := chromedp.Run(newCtx, elementScreenshot(linkURL, &screenshot)); err != nil {
		// Prefer context cause when our timeout context was canceled for clearer diagnostics
		if screenshotCtx.Err() != nil {
			if cause := context.Cause(screenshotCtx); cause != nil {
				err = cause
			}
		}
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
		chromedp.WaitVisible("body", chromedp.ByQuery), // Wait for page to load
		chromedp.Sleep(1 * time.Second),                // Additional wait for dynamic content
		chromedp.CaptureScreenshot(res),
	}
}
