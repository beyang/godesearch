package godesearch

import (
	"fmt"
	"github.com/google/go-github/github"
)

type Searcher struct {
	Client *github.Client
}

func (s *Searcher) Search(query string) {
	opts := &github.SearchOptions{
		PerPage: 100,
		Page:    1,
	}
	result, _, err := s.Client.Search.Code(query, opts)
	fmt.Printf("Result: %+v", result)
}
