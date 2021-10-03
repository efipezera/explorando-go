package gerador

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//Título obtém o título de uma página html.
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		//erro: panic: runtime error: index out of range [1] with length 0
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile(`<title>(.*?)<\\/title>`)
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}
