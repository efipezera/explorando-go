package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

//Título obtém o título de uma página html.
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		//panic: runtime error: index out of range [1] with length 0
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile(`<title>(.*?)<\\/title>`)
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func oMaisRapido(url1, url2, url3 string) string {
	c1 := Titulo(url1)
	c2 := Titulo(url2)
	c3 := Titulo(url3)

	//Estrutura de controle própria para concorrência.
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(time.Millisecond * 1000):
		return "Todos perderam!"
		//default:
		//	return "Sem resposta..."
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.udemy.com/",
		"https://www.youtube.com/",
		"https://www.instagram.com/",
	)
	fmt.Println(campeao)
}
