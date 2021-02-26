package main

import (
	"fmt"
	"net/http"
	"sync"
)

func returnTyoe(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s'\n", err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s: %s\n", url, ctype)
}

func introConcur() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://linkedin.com",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnTyoe(url)
			wg.Done()
		}(url)
		// returnTyoe(url)
	}
	wg.Wait()
}
