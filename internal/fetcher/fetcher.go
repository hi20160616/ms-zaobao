package fetcher

import (
	"context"
	"errors"
	"log"
	"os"
	"sort"

	"github.com/hi20160616/ms-zaobao/configs"
)

// Fetch fetch and storage all stuffs to `db/articles.json`
func Fetch() error {
	defer log.Printf("[%s] Fetch Done.", configs.Data.MS.Title)
	log.Printf("[%s] Fetching ...", configs.Data.MS.Title)

	as, err := fetch(context.Background())
	if err != nil {
		return err
	}

	as, err = merge(as)
	if err != nil {
		return err
	}

	sort.Sort(sort.Reverse(ByUpdateTime(as)))

	return storage(as)
}

// fetch fetch all articles by url set in config.json
func fetch(ctx context.Context) (as []*Article, err error) {
	links, err := fetchLinks()
	if err != nil {
		return
	}
	for _, link := range links {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			a := NewArticle()
			a, err = a.fetchArticle(link)
			if err != nil {
				if errors.Is(err, ErrTimeOverDays) {
					continue
				}
				log.Printf("[%s] fetch error: %v, link: %s",
					configs.Data.MS.Title, err, link)
			}
			// ignore redundant articles
			exist := false
			for _, _a := range as {
				if a.Title == _a.Title || a.Id == _a.Id {
					exist = true
				}
			}
			if !exist {
				as = append(as, a)
			}
		}
	}
	return
}

// merge will merge local data and fetched data from db/articles.json and website respectively
func merge(as []*Article) ([]*Article, error) {
	dbAs, err := load()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return as, nil
		}
		return nil, err
	}
	as = append(as, dbAs...)
	return as, nil
}
