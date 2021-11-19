package fetcher

import (
	"context"
	"testing"

	"github.com/hi20160616/ms-zaobao/configs"
)

func TestFetch(t *testing.T) {
	if err := configs.Reset("../../"); err != nil {
		t.Error(err)
	}

	if err := Fetch(context.Background()); err != nil {
		t.Error(err)
	}
}
