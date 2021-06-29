package fetcher

import (
	"fmt"
	"testing"
)

func TestGetLinks(t *testing.T) {
	links, err := getLinks("https://www.zaobao.com/realtime/world")
	if err != nil {
		t.Error(err)
	}
	for _, link := range links {
		fmt.Println(link)
	}
}
