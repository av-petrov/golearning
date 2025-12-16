package main

import (
	"testing"
)

func Test_Crawl(t *testing.T) {
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
	if len(cache.seen) != 5 {
		t.Error("incorrect result: expected 5, got:", len(cache.seen))
	}
}
