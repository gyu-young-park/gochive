package worker

import (
	"fmt"
	"github/gyu-young-park/go-archive/repository"

	"github.com/gocolly/colly/v2"
)

type mediumWork struct {
	crawler  *colly.Collector
	postList []repository.Post
}

func newMediumWork() *mediumWork {
	m := &mediumWork{
		crawler: colly.NewCollector(),
	}
	m.ready()
	return m
}

func (m *mediumWork) ready() {
	fmt.Println("hello workder")
	m.crawler.OnHTML("article", func(e *colly.HTMLElement) {
		post := repository.Post{}
		post.Link = e.ChildAttrs("a", "href")[2]
		post.Title = e.ChildText("h2")
		m.postList = append(m.postList, post)
	})
}

func (m *mediumWork) do() {
	m.crawler.Visit("https://medium.com/tag/go/latest")
}
