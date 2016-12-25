package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	totalResult := 0
	// Задаем общее количество сайтов
	var allSite int
	fmt.Print("How many sites: ")
	fmt.Scanln(&allSite)
	// Заполняем срез сайтами
	urls := make([]string, allSite)
	ch := make(chan int, allSite)
	for i := 0; i < allSite; i++ {
		fmt.Print("Enter site ", i+1, ": ")
		fmt.Scanln(&urls[i])
	}
	// канал для подсчета Total

	// обрабатываем сразу несколько сайтов
	for s := range urls {
		go countGo(urls, s, ch)

		// общее количество найденных вхождений
		for g := 0; g < allSite; g++ {
			totalResult += <-ch
		}

		fmt.Printf("Total: %d\n", totalResult)
	}
}

func countGo(urls []string, s int, ch chan int) {
	searchGo := "Go"

	url := urls[s]
	resp, err := http.Get(url)
	//er(err)
	if err != nil {
		return
	}
	site, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	text := string(site)
	count := strings.Count(text, searchGo)

	fmt.Printf("Count for %s = %d\n", url, count)
	ch <- count
}
