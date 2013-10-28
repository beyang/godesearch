package main

import (
	"flag"
	"fmt"
	gs "github.com/beyang/godesearch/godesearch"
)

func main() {
	flag.Parse()
	query := flag.Arg(0)

	s := gs.NewSearcher("")
	// result, err := s.Search(*query)
	// if err != nil {
	// 	fmt.Printf("Error in search: %s", err)
	// }
	// fmt.Printf("Result: %+v", result)

	_, err := s.Search(query)
	if err != nil {
		fmt.Printf("Error in search: %s", err)
	}
}
