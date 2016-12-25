package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	const k = 5
	totalResult := 0
	searchGo := "Go"
	var allSite int
	fmt.Print("How many sites: ")
	fmt.Scanln(&allSite)
	urls := make([]string, allSite)
	for i := 0; i < allSite; i++ {
		fmt.Print("Enter site ", i+1, ": ")
		fmt.Scanln(&urls[i])
	}

	ch := make(chan int, allSite)

	for s := 0; s < allSite; s++ {
		go func(s int) {
			url := urls[s]
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			site, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			text := string(site)
			count := strings.Count(text, searchGo)

			fmt.Printf("Count for %s = %d\n", url, count)
			ch <- count
		}(s)
	}

	for g := 0; g < allSite; g++ {
		z := <-ch
		totalResult += z
	}

	fmt.Printf("Total: %d\n", totalResult)

}
