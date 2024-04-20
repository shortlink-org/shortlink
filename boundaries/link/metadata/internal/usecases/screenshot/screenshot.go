/*
Metadata Service. Application layer
*/

package screenshot

import (
	"context"
	"net/url"

	"github.com/chromedp/chromedp"

	s3Repository "github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/infrastructure/repository/media"
)

func New(ctx context.Context, media *s3Repository.Service) (*UC, error) {
	return &UC{
		media: media,
	}, nil
}

func (s *UC) Get(ctx context.Context, linkURL string) (*url.URL, error) {
	// return s.media.Get(ctx, linkURL)
	panic("not implemented")
}

func (s *UC) Set(ctx context.Context, linkURL string) error {
	chromedp.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4585.0 Safari/537.36")

	newCtx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	// capture screenshot of an element
	var screenshot []byte
	err := chromedp.Run(newCtx, elementScreenshot(linkURL, &screenshot))
	if err != nil {
		return err
	}

	// err = s.media.Put(ctx, fmt.Sprintf("%s.png", linkURL), screenshot)
	// if err != nil {
	// 	return err
	// }

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
