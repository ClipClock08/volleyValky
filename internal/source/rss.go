package source

import (
	"context"
	"github.com/SlyMarbo/rss"
	"news-feed-bot/internal/model"
)

type RSSSource struct {
	URL        string
	ID         int64
	SourceName string
}

func NewRssSourceFromModel(m model.Source) RSSSource {
	return RSSSource{
		URL:        m.FeedURL,
		ID:         m.ID,
		SourceName: m.Name,
	}
}

func (s RSSSource) loadFeed(ctx context.Context, url string) (*rss.Feed, error) {
	var (
		feedCh = make(chan *rss.Feed)
		errCh  = make(chan error)
	)
	go func() {
		feed, err := rss.Fetch(url)
		if err != nil {
			errCh <- err
			return
		}
		feedCh <- feed
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errCh:
		return nil, err
	case feed := <-feedCh:
		return feed, nil
	}
}
