package main

import (
	"fmt"

	latest "github.com/tcnksm/go-latest"
)

func main() {
	githubTag := &latest.GithubTag{
		Owner:      "ngratin",
		Repository: "go-fmt-test",
	}

	res, _ := latest.Check(githubTag, "0.1.0")
	if res.Outdated {
		fmt.Println(res.Current)
		fmt.Println("outdate")
	}

	fmt.Println("hello")
	fmt.Println("hello")
}
