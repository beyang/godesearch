package godesearch

import (
	"fmt"
	"github.com/google/go-github/github"
)

type Searcher struct {
	Client *github.Client
}

type SearchResult struct {
}

func NewSearcher(accessToken string) *Searcher {
	return &Searcher{
		Client: github.NewClient(nil),
	}
}

func (s *Searcher) Search(query string) (*SearchResult, error) {
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
			Page:    1,
		},
	}
	res, _, err := s.Client.Search.Code(query, opts)
	if err != nil {
		return nil, err
	}

	fmt.Printf("res: %+v\n", res)
	// TODO

	return &SearchResult{}, nil
}
