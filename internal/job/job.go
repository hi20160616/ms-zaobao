package job

import (
	"context"
	"log"
	"time"

	"github.com/hi20160616/ms-zaobao/configs"
	"github.com/hi20160616/ms-zaobao/internal/fetcher"
	"github.com/pkg/errors"
)

func Crawl(ctx context.Context) error {
	t, err := time.ParseDuration(configs.Data.MS.Heartbeat)
	if err != nil {
		return err
	}
	for {
		select {
		case <-time.Tick(t):
			if err := fetcher.Fetch(); err != nil {
				if !errors.Is(err, fetcher.ErrTimeOverDays) {
					log.Printf("%#v", err)
				}
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// Stop is nil now
func Stop(ctx context.Context) error {
	log.Println("Job gracefully stopping.")
	// return error can define here, so it will display on frontend
	return ctx.Err()
}
